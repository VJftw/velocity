module Page.Project.Commit.Overview exposing (Model, Msg(..), initialModel, maybeBuildFromTask, taskBuilds, update, view, viewCommitDetails, viewTaskList, viewTaskListItem)

import Data.Build as Build exposing (Build)
import Data.Commit as Commit exposing (Commit)
import Data.Project as Project exposing (Project)
import Data.Session as Session exposing (Session)
import Data.Task as ProjectTask
import Html exposing (..)
import Html.Attributes exposing (..)
import Navigation
import Page.Helpers exposing (formatDateTime)
import Page.Project.Commit.Route as CommitRoute
import Page.Project.Route as ProjectRoute
import Route
import Util exposing ((=>))
import Views.Build exposing (viewBuildStatusIcon, viewBuildTextClass)
import Views.Commit exposing (branchList, infoPanel, truncateCommitMessage)
import Views.Helpers exposing (onClickPage)


-- MODEL --


type alias Model =
    {}


initialModel : Model
initialModel =
    {}



-- VIEW --


view : Project -> Commit -> List ProjectTask.Task -> List Build -> Html Msg
view project commit tasks builds =
    div []
        [ viewCommitDetails commit
        , hr [] []
        , viewTaskList project commit tasks builds
        ]


viewCommitDetails : Commit -> Html Msg
viewCommitDetails commit =
    div [ class "mb-2" ]
        [ div [ class "" ]
            [ infoPanel commit
            , branchList commit
            ]
        ]


viewTaskList : Project -> Commit -> List ProjectTask.Task -> List Build -> Html Msg
viewTaskList project commit tasks builds =
    let
        taskList =
            tasks
                |> List.filter (.name >> ProjectTask.nameToString >> String.isEmpty >> not)
                |> List.sortBy (.name >> ProjectTask.nameToString)
                |> List.map (viewTaskListItem project commit builds)
                |> div [ class "list-group list-group-flush" ]
    in
        div [ class "card my-2" ]
            [ taskList
            ]


taskBuilds : ProjectTask.Task -> List Build -> List Build
taskBuilds task builds =
    builds
        |> List.filter (\b -> ProjectTask.idEquals task.id b.task.id)


maybeBuildFromTask : ProjectTask.Task -> List Build -> Maybe Build
maybeBuildFromTask task builds =
    taskBuilds task builds
        |> List.head


viewTaskListItem : Project -> Commit -> List Build -> ProjectTask.Task -> Html Msg
viewTaskListItem project commit builds task =
    let
        route =
            CommitRoute.Task task.slug Nothing
                |> ProjectRoute.Commit commit.hash
                |> Route.Project project.slug

        maybeBuild =
            maybeBuildFromTask task builds

        textClass =
            maybeBuild
                |> Maybe.map viewBuildTextClass
                |> Maybe.withDefault ""
    in
        a
            [ class (textClass ++ " list-group-item list-group-item-action flex-column align-items-center justify-content-between")
            , Route.href route
            , onClickPage NewUrl route
            ]
            [ div [ class "" ] [ h6 [ class "mb-1" ] [ text (ProjectTask.nameToString task.name) ] ]
            , p [ class "mb-1" ] [ text task.description ]
            ]



-- UPDATE --


type Msg
    = NewUrl String


update : Project -> Session msg -> Msg -> Model -> ( Model, Cmd Msg )
update project session msg model =
    case msg of
        NewUrl url ->
            model => Navigation.newUrl url
