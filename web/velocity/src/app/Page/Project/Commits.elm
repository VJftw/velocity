module Page.Project.Commits exposing (..)

import Context exposing (Context)
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, on, targetValue)
import Data.Commit as Commit exposing (Commit)
import Data.Session as Session exposing (Session)
import Data.Project as Project exposing (Project)
import Data.Branch as Branch exposing (Branch)
import Data.PaginatedList as PaginatedList exposing (PaginatedList, Paginated(..))
import Page.Errored as Errored exposing (PageLoadError, pageLoadError)
import Page.Helpers exposing (formatDate, formatTime, sortByDatetime)
import Data.AuthToken as AuthToken exposing (AuthToken)
import Request.Project
import Request.Commit
import Request.Errors
import Util exposing ((=>))
import Task exposing (Task)
import Views.Page as Page
import Http
import Dict exposing (Dict)
import Time.DateTime as DateTime exposing (DateTime)
import Time.Date as Date exposing (Date)
import Page.Helpers exposing (formatDate)
import Route exposing (Route)
import Page.Project.Route as ProjectRoute
import Page.Project.Commit.Route as CommitRoute
import Navigation
import Views.Helpers exposing (onClickPage)
import Json.Encode as Encode
import Component.DropdownFilter as DropdownFilter
import Dom
import Bootstrap.Button as Button


-- MODEL --


type alias Model =
    { commits : List Commit
    , total : Int
    , page : Int
    , submitting : Bool
    , branch : Maybe Branch
    , dropdownState : DropdownFilter.DropdownState
    , branchFilterTerm : String
    }


loadCommits : Context -> Project.Slug -> Maybe AuthToken -> Maybe Branch -> Int -> Task Request.Errors.HttpError (PaginatedList Commit)
loadCommits context projectSlug maybeAuthToken maybeBranch page =
    maybeAuthToken
        |> Request.Commit.list context projectSlug (Maybe.map .name maybeBranch) perPage page


init : Context -> Session msg -> List Branch -> Project.Slug -> Maybe Branch.Name -> Maybe Int -> Task PageLoadError Model
init context session branches projectSlug maybeBranchName maybePage =
    let
        defaultPage =
            Maybe.withDefault 1 maybePage

        maybeAuthToken =
            Maybe.map .token session.user

        maybeBranch =
            branches
                |> List.filter (\b -> maybeBranchName == Just b.name)
                |> List.head

        initialModel (Paginated { results, total }) =
            { commits = results
            , total = total
            , page = defaultPage
            , submitting = False
            , branch = maybeBranch
            , dropdownState = DropdownFilter.initialDropdownState
            , branchFilterTerm = ""
            }

        handleLoadError _ =
            pageLoadError Page.Project "Project unavailable."
    in
        Task.map initialModel (loadCommits context projectSlug maybeAuthToken maybeBranch defaultPage)
            |> Task.mapError handleLoadError


perPage : Int
perPage =
    10


branchFilterConfig : DropdownFilter.Config Msg Branch
branchFilterConfig =
    { dropdownMsg = BranchFilterDropdownMsg
    , termMsg = BranchFilterTermMsg
    , noOpMsg = NoOp
    , selectItemMsg = FilterBranch
    , labelFn = (.name >> Just >> Branch.nameToString)
    }


branchFilterContext : List Branch -> Model -> DropdownFilter.Context Branch
branchFilterContext branches { dropdownState, branchFilterTerm, branch } =
    let
        items =
            branches
                |> List.filter .active
                |> List.sortBy (.name >> Just >> Branch.nameToString)
                |> List.sortBy (branchFilterConfig.labelFn)
    in
        { items = items
        , dropdownState = dropdownState
        , filterTerm = branchFilterTerm
        , selectedItem = branch
        }



-- SUBSCRIPTIONS --


subscriptions : List Branch -> Model -> Sub Msg
subscriptions branches model =
    branchFilterContext branches model
        |> DropdownFilter.subscriptions branchFilterConfig



-- CHANNELS --


channelName : Project.Slug -> String
channelName projectSlug =
    "project:" ++ (Project.slugToString projectSlug)


events : Project.Slug -> Dict String (List ( String, Encode.Value -> Msg ))
events projectSlug =
    let
        pageEvents =
            [ ( "commit:new", always RefreshCommitList )
            , ( "commit:update", always RefreshCommitList )
            , ( "commit:deleted", always RefreshCommitList )
            ]
    in
        Dict.singleton (channelName projectSlug) (pageEvents)



-- VIEW --


view : Project -> List Branch -> Model -> Html Msg
view project branches model =
    let
        commits =
            commitListToDict model.commits
                |> viewCommitListContainer project

        branchFilter =
            branchFilterContext branches model
                |> DropdownFilter.view branchFilterConfig

        refreshCommitsButton =
            refreshButton project model

        buttons =
            div [ class "btn-toolbar" ]
                [ branchFilter
                , refreshCommitsButton
                ]

        paginationToolbar =
            div [ class "btn-toolbar" ]
                [ pagination model.page model.total project model.branch ]
    in
        div []
            [ h4 [ class "mb-2" ] [ text "Commits" ]
            , buttons
            , commits
            , paginationToolbar
            ]


commitListToDict : List Commit -> Dict ( Int, Int, Int ) (List Commit)
commitListToDict commits =
    let
        reducer commit dict =
            let
                date =
                    commit.date
                        |> DateTime.date
                        |> Date.toTuple

                insert =
                    case Dict.get date dict of
                        Just exists ->
                            commit :: exists

                        Nothing ->
                            [ commit ]
            in
                Dict.insert date insert dict
    in
        List.foldl reducer Dict.empty commits


viewCommitListContainer : Project -> Dict ( Int, Int, Int ) (List Commit) -> Html Msg
viewCommitListContainer project dict =
    let
        listItemToDate dateListItem =
            dateListItem
                |> Tuple.first
                |> Date.fromTuple

        sortDateList a b =
            listItemToDate a
                |> Date.compare (listItemToDate b)
    in
        dict
            |> Dict.toList
            |> List.sortWith sortDateList
            |> List.take perPage
            |> List.map (viewCommitList project)
            |> div [ class "mt-3" ]


viewCommitList : Project -> ( ( Int, Int, Int ), List Commit ) -> Html Msg
viewCommitList project ( dateTuple, commits ) =
    let
        commitListItems =
            sortByDatetime .date commits
                |> List.map (viewCommitListItem project.slug)

        formattedDate =
            Date.fromTuple dateTuple
                |> formatDate
    in
        div []
            [ h6 [ class "mb-2 mt-2 text-muted" ] [ text formattedDate ]
            , div [ class "card" ]
                [ div [ class "list-group list-group-flush" ] commitListItems
                ]
            ]


viewCommitListItem : Project.Slug -> Commit -> Html Msg
viewCommitListItem slug commit =
    let
        truncatedHash =
            Commit.truncateHash commit.hash

        route =
            Route.Project slug <| ProjectRoute.Commit commit.hash CommitRoute.Overview

        branchList =
            commit.branches
                |> List.map (\b -> span [ class "badge badge-secondary" ] [ i [ class "fa fa-code-fork" ] [], text (" " ++ (Branch.nameToString (Just b))) ])
                |> List.map (\b -> li [ class "list-inline-item" ] [ b ])
                |> (ul [ class "mb-0 list-inline" ])
    in
        a [ class "list-group-item list-group-item-action flex-column align-items-start", Route.href route, onClickPage NewUrl route ]
            [ div [ class "d-flex w-100 justify-content-between" ]
                [ h6 [ class "mb-1 text-overflow" ] [ text commit.message ]
                , small [] [ text truncatedHash ]
                ]
            , div [ class "d-flex w-100 justify-content-between" ]
                [ small [] [ strong [] [ text commit.author ], text (" commited at " ++ formatTime commit.date) ]
                , branchList
                ]
            ]


breadcrumb : Project -> List ( Route, String )
breadcrumb project =
    [ ( Route.Project project.slug (ProjectRoute.Commits Nothing Nothing), "Commits" ) ]


viewBreadcrumbExtraItems : Project -> Model -> Html Msg
viewBreadcrumbExtraItems project model =
    text ""


refreshButton : Project -> Model -> Html Msg
refreshButton project model =
    let
        submitting =
            project.synchronising || model.submitting

        iconClassList =
            [ ("fa fa-refresh" => True)
            , ("fa-spin fa-fw" => submitting)
            ]
    in
        Button.button
            [ Button.outlineSecondary
            , Button.attrs
                [ class "ml-auto btn btn-dark"
                , onClick SubmitSync
                , disabled submitting
                ]
            ]
            [ i [ classList iconClassList ] [] ]


pagination : Int -> Int -> Project -> Maybe Branch -> Html Msg
pagination activePage total project maybeBranch =
    let
        totalPages =
            ceiling (toFloat total / toFloat perPage)
    in
        if totalPages > 1 then
            List.range 1 totalPages
                |> List.map (\page -> pageLink page (page == activePage) project maybeBranch)
                |> ul [ class "pagination" ]
        else
            Html.text ""


pageLink : Int -> Bool -> Project -> Maybe Branch -> Html Msg
pageLink page isActive project maybeBranch =
    let
        route =
            Route.Project project.slug <| ProjectRoute.Commits (Maybe.map .name maybeBranch) (Just page)
    in
        li [ classList [ "page-item" => True, "active" => isActive ] ]
            [ a
                [ class "page-link"
                , Route.href route
                , onClickPage NewUrl route
                ]
                [ text (toString page) ]
            ]



-- UPDATE --


type Msg
    = SubmitSync
    | SyncCompleted (Result Request.Errors.HttpError Project)
    | FilterBranch (Maybe Branch)
    | SelectPage Int
    | NewUrl String
    | RefreshCommitList
    | RefreshCompleted (Result Request.Errors.HttpError (PaginatedList Commit))
    | BranchFilterDropdownMsg DropdownFilter.DropdownState
    | BranchFilterTermMsg String
    | NoOp


update : Context -> Project -> Session msg -> Msg -> Model -> ( Model, Cmd Msg )
update context project session msg model =
    case msg of
        NewUrl newUrl ->
            model => Navigation.newUrl newUrl

        SubmitSync ->
            let
                cmdFromAuth authToken =
                    authToken
                        |> Request.Project.sync context project.slug
                        |> Task.attempt SyncCompleted

                cmd =
                    session
                        |> Session.attempt "sync project" cmdFromAuth
                        |> Tuple.second
            in
                { model | submitting = True } => cmd

        SyncCompleted (Ok _) ->
            { model | submitting = False }
                => Cmd.none

        SyncCompleted (Err _) ->
            { model | submitting = False }
                => Cmd.none

        SelectPage page ->
            let
                uriEncoded =
                    model.branch
                        |> Maybe.map .name
                        |> Maybe.andThen
                            (\(Branch.Name slug) ->
                                slug
                                    |> Http.encodeUri
                                    |> Branch.Name
                                    |> Just
                            )

                newRoute =
                    Route.Project project.slug <| ProjectRoute.Commits uriEncoded (Just page)
            in
                model => Route.modifyUrl newRoute

        FilterBranch maybeBranch ->
            let
                uriEncoded =
                    maybeBranch
                        |> Maybe.map .name
                        |> Maybe.andThen
                            (\(Branch.Name slug) ->
                                slug
                                    |> Http.encodeUri
                                    |> Branch.Name
                                    |> Just
                            )

                newRoute =
                    Route.Project project.slug <| ProjectRoute.Commits uriEncoded (Just 1)
            in
                { model | commits = [] }
                    => Route.modifyUrl newRoute

        RefreshCommitList ->
            let
                refreshTask =
                    loadCommits context project.slug (Maybe.map .token session.user) model.branch model.page
            in
                model => Task.attempt RefreshCompleted refreshTask

        RefreshCompleted (Ok (Paginated { results, total })) ->
            { model
                | commits = results
                , total = total
            }
                => Cmd.none

        RefreshCompleted (Err _) ->
            model => Cmd.none

        BranchFilterDropdownMsg state ->
            { model | dropdownState = state }
                => Task.attempt (always NoOp) (Dom.focus "filter-branch-input")

        BranchFilterTermMsg term ->
            { model | branchFilterTerm = term }
                => Cmd.none

        NoOp ->
            model => Cmd.none
