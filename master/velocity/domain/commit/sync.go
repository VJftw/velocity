package commit

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/velocity-ci/velocity/master/velocity/domain"
	"github.com/velocity-ci/velocity/master/velocity/domain/project"
	"github.com/velocity-ci/velocity/master/velocity/domain/task"
	git "gopkg.in/src-d/go-git.v4"
)

func sync(p *domain.Project, projectManager *project.BoltManager, commitManager *BoltManager) {
	repo, dir, err := project.Clone(p.Name, p.Repository, p.PrivateKey, false)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer os.RemoveAll(dir) // clean up

	refIter, err := repo.References()
	if err != nil {
		log.Fatal(err)
		return
	}
	w, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		r, err := refIter.Next()
		if err != nil {
			break
		}

		fmt.Println(r)
		commit, err := repo.CommitObject(r.Hash())

		if err != nil {
			break
		}

		mParts := strings.Split(commit.Message, "-----END PGP SIGNATURE-----")
		message := mParts[0]
		if len(mParts) > 1 {
			message = mParts[1]
		}

		branch := strings.Join(strings.Split(r.Name().Short(), "/")[1:], "/")

		c := domain.Commit{
			Branch:  branch,
			Hash:    commit.Hash.String(),
			Message: strings.TrimSpace(message),
			Author:  commit.Author.Email,
			Date:    commit.Committer.When,
		}

		commitManager.SaveCommitForProject(p, &c)

		err = w.Checkout(&git.CheckoutOptions{
			Hash: commit.Hash,
		})

		if err != nil {
			fmt.Println(err)
		}

		SHA := r.Hash().String()
		shortSHA := SHA[:7]
		describe := shortSHA

		gitParams := map[string]task.Parameter{
			"GIT_SHA": task.Parameter{
				Value: SHA,
			},
			"GIT_SHORT_SHA": task.Parameter{
				Value: shortSHA,
			},
			"GIT_BRANCH": task.Parameter{
				Value: branch,
			},
			"GIT_DESCRIBE": task.Parameter{
				Value: describe,
			},
		}

		if _, err := os.Stat(fmt.Sprintf("%s/tasks/", dir)); err == nil {
			filepath.Walk(fmt.Sprintf("%s/tasks/", dir), func(path string, f os.FileInfo, err error) error {
				if !f.IsDir() && strings.HasSuffix(f.Name(), ".yml") || strings.HasSuffix(f.Name(), ".yaml") {
					taskYml, _ := ioutil.ReadFile(fmt.Sprintf("%s/tasks/%s", dir, f.Name()))
					t := task.ResolveTaskFromYAML(string(taskYml), gitParams)
					commitManager.SaveTaskForCommitInProject(&t, &c, p)
				}
				return nil
			})
		}
	}

	p.UpdatedAt = time.Now()
	p.Synchronising = false
	projectManager.Save(p)

}
