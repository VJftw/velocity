defmodule Architect.Projects do
  @moduledoc """
  The Projects context.
  """

  import Ecto.Query, warn: false
  alias Architect.Repo
  alias Architect.Projects.{Repository, Project, Starter}
  alias Architect.Events

  alias Architect.Accounts.User
  use Supervisor
  require Logger

  @registry __MODULE__.Registry
  @supervisor __MODULE__.Supervisor

  def start_link(_opts \\ []) do
    Logger.debug("Starting #{Atom.to_string(__MODULE__)}")
    Supervisor.start_link(__MODULE__, :ok, name: __MODULE__)
  end

  @doc """
  Returns the list of projects.

  ## Examples

      iex> list_projects()
      [%Project{}, ...]

  """
  def list_projects() do
    Repo.all(Project)
  end

  @doc """
  Gets a single project by id.

  Raises `Ecto.NoResultsError` if the Known host does not exist.
  """
  def get_project!(id), do: Repo.get!(Project, id)

  @doc """
  Gets a single project by slug.

  Raises `Ecto.NoResultsError` if the Project does not exist.

  ## Examples

      iex> get_project_by_slug!("velocity")
      %KnownHost{}

      iex> get_project_by_slug!("Not a slug")
      ** (Ecto.NoResultsError)

  """
  def get_project_by_slug!(slug), do: Repo.get_by!(Project, slug: slug)

  @doc """
  Creates a project.

  ## Examples

      iex> create_project(%User{}, "https://github.com/velocity-ci/velocity.git")
      {:ok, %Project{}, %Event{}}

      iex> create_project(%User{}, "banter)
      {:error, %Ecto.Changeset{}}

  """
  def create_project(%User{} = u, address) when is_binary(address) do
    Repo.transaction(fn ->
      changeset = Project.changeset(%Project{}, %{address: address, created_by_id: u.id})

      case Repo.insert(changeset) do
        {:ok, p} ->
          event = Events.create_event!(u, p, %{type: :project_created})

          {p, event}

        {:error, e} ->
          Repo.rollback(e)
      end
    end)
  end

  @doc ~S"""
  Get a list of branches

  ## Examples

      iex> list_branches(project)
      [%Branch{}, ...]

  """
  def list_branches(%Project{} = project),
    do: call_repository(project, {:list_branches, []})

  @doc ~S"""
  Get a list of branches for a specific commit SHA

  ## Examples

      iex> list_branches_for_commit(project, "925fbc450c8bdb7665ec3af3129ce715927433fe")
      [%Branch{}, ...]

  """
  def list_branches_for_commit(%Project{} = project, sha) when is_binary(sha),
    do: call_repository(project, {:list_branches_for_commit, [sha]})

  @doc ~S"""
  Get a list of commits by branch


  ## Examples

      iex> list_commits(project, "master")
      [%Commit{}, ...]

  """
  def list_commits(%Project{} = project, branch) when is_binary(branch),
    do: call_repository(project, {:list_commits, [branch]})

  @doc ~S"""
  Get the default branch

  ## Examples

      iex> default_branch(project)
      %Branch{}

  """
  def default_branch(%Project{} = project),
    do: call_repository(project, {:default_branch, []})

  @doc ~S"""
  Get specific branch

  ## Examples

      iex> get_branch(project, "master")
      %Branch{}

  """
  def get_branch(%Project{} = project, branch) when is_binary(branch),
    do: call_repository(project, {:get_branch, [branch]})

  @doc ~S"""
  Get the amount of commits for the project

  ## Examples

      iex> commit_count(project)
      123

  """
  def commit_count(%Project{} = project),
    do: call_repository(project, {:commit_count, []})

  @doc ~S"""
  Get the amount of commits for the project, for a specific branch

  ## Examples

      iex> commit_count(project, "master")
      42

  """
  def commit_count(%Project{} = project, branch) when is_binary(branch),
    do: call_repository(project, {:commit_count, [branch]})

  @doc ~S"""
  List Blueprint

  ## Examples

      iex> list_blueprints(project, {:sha, "925fbc450c8bdb7665ec3af3129ce715927433fe"})
      [%Architect.Projects.Blueprint{}, ...]

  """
  def list_blueprints(%Project{} = project, selector),
    do: call_repository(project, {:list_blueprints, [selector]})

  @doc ~S"""
  Project Configuration

  """
  def project_configuration(%Project{} = project),
    do: call_repository(project, {:project_configuration, []})

  @doc ~S"""
  Get the construction plan for a Blueprint on a commit sha
  """
  def plan_blueprint(%Project{} = project, branch_name, commit, blueprint_name),
    do: call_repository(project, {:plan_blueprint, [branch_name, commit, blueprint_name]}, false)

  ### Server

  @impl true
  def init(:ok) do
    :ets.new(:simple_cache, [:named_table, :public])

    children = [
      {Registry, keys: :unique, name: @registry},
      {DynamicSupervisor, name: @supervisor, strategy: :one_for_one, max_restarts: 3},
      worker(
        Starter,
        [%{registry: @registry, supervisor: @supervisor, projects: list_projects()}],
        restart: :transient
      )
    ]

    Logger.info("Running #{Atom.to_string(__MODULE__)}")

    Supervisor.init(children, strategy: :one_for_one)
  end

  @doc false
  defp call_repository(project, callback, cache \\ true, attempt \\ 1)

  defp call_repository(_, _, _, attempt) when attempt > 2, do: {:error, "Failed"}

  defp call_repository(
         %Project{address: address, name: name} = project,
         {fun, args},
         cache,
         attempt
       ) do
    case Registry.lookup(@registry, "#{address}-#{name}") do
      [{repository, _}] ->
        try do
          Architect.ETSCache.get(Repository, cache, fun, [repository | args])
        catch
          kind, reason ->
            Logger.warn(
              "Failed to call repository #{address} #{name}: #{inspect(fun)} #{inspect(args)} (#{
                inspect(kind)
              }-#{inspect(reason)}), #{inspect(attempt)}..."
            )

            Process.sleep(1_000)

            call_repository(project, {fun, args}, attempt + 1, cache)
        end

      [] ->
        Logger.warn(
          "Failed to call builder #{address} #{name} on #{inspect(@registry)}; does not exist"
        )

        call_repository(project, {fun, args}, attempt + 1, cache)
    end
  end
end

defmodule Architect.Projects.Starter do
  use Task
  require Logger
  alias Architect.Projects.Repository

  def start_link(opts) do
    Task.start_link(__MODULE__, :run, [opts])
  end

  def run(%{projects: projects, supervisor: _supervisor, registry: _registry}) do
    for project <- projects do
      repository_name =
        {:via, Registry, {Architect.Projects.Registry, "#{project.address}-#{project.name}"}}

      {:ok, _pid} =
        DynamicSupervisor.start_child(
          # supervisor,
          Architect.Projects.Supervisor,
          {Repository, {project.address, project.private_key, repository_name}}
        )
    end
  end
end
