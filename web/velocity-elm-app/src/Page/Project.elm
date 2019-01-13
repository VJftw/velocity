module Page.Project exposing (Model, Msg, init, subscriptions, toContext, toSession, update, view)

import Asset
import Browser.Events
import Context exposing (Context)
import Element exposing (..)
import Element.Background as Background
import Element.Border as Border
import Element.Button as Button
import Element.Events exposing (onClick)
import Element.Font as Font
import Form.Input as Input
import Form.Select as Select
import Http
import Icon
import Json.Decode as Decode
import Loading
import PaginatedList exposing (PaginatedList)
import Palette
import Project exposing (Project)
import Project.Branch as Branch exposing (Branch)
import Project.Branch.Name as BranchName
import Project.Build as Build exposing (Build)
import Project.Commit as Commit exposing (Commit)
import Project.Id exposing (Id)
import Project.Slug exposing (Slug)
import Project.Task as Task exposing (Task)
import Project.Task.Name as TaskName
import Session exposing (Session)
import Task as BaseTask
import Time
import Timestamp


-- Model


type alias Model msg =
    { session : Session msg
    , context : Context msg
    , tasks : Status (List Task)
    , timeZone : Time.Zone
    , slug : Slug
    , branchDropdown : BranchDropdown
    , builds : Status (PaginatedList Build)
    }


type BranchDropdown
    = BranchDropdown BranchDropdownStatus


type BranchDropdownStatus
    = Open
    | ListenClicks
    | Closed


type Status a
    = Loading
    | LoadingSlowly
    | Loaded a
    | Failed


init : Session msg -> Context msg -> Slug -> ( Model msg, Cmd Msg )
init session context projectSlug =
    let
        maybeProject =
            Session.projectWithSlug projectSlug session

        maybeCred =
            Session.cred session

        baseUrl =
            Context.baseUrl context

        ( builds, buildsRequest ) =
            case ( maybeProject, maybeCred ) of
                ( Just project, Just cred ) ->
                    ( Loading
                    , Cmd.none
                      --                    , Build.list cred baseUrl projectSlug CompletedLoadBuilds
                    )

                _ ->
                    ( Failed, Cmd.none )

        ( tasks, tasksRequest ) =
            case ( maybeProject, maybeCred ) of
                ( Just project, Just cred ) ->
                    ( Loading
                    , BaseTask.attempt CompletedLoadTasks <|
                        Task.byBranch cred baseUrl projectSlug BranchName.default
                    )

                _ ->
                    ( Failed, Cmd.none )
    in
        ( { session = session
          , context = context
          , tasks = tasks
          , timeZone = Time.utc
          , slug = projectSlug
          , branchDropdown = BranchDropdown Closed
          , builds = Loading
          }
        , Cmd.batch
            [ buildsRequest
            , tasksRequest
            , BaseTask.perform (\_ -> PassedSlowLoadThreshold) Loading.slowThreshold
            , BaseTask.perform GotTimeZone Time.here
            ]
        )



-- Subscriptions


subscriptions : Model msg -> Sub Msg
subscriptions { branchDropdown } =
    case branchDropdown of
        BranchDropdown Open ->
            Browser.Events.onAnimationFrame
                (\_ -> BranchDropdownListenClicks)

        BranchDropdown ListenClicks ->
            Browser.Events.onClick
                (Decode.succeed BranchDropdownClose)

        BranchDropdown Closed ->
            Sub.none



-- Update


type Msg
    = NoOp
    | BranchDropdownToggleClicked
    | BranchDropdownListenClicks
    | BranchDropdownClose
    | CompletedLoadBuilds (Result Http.Error (PaginatedList Build))
    | CompletedLoadTasks (Result Http.Error (List Task))
    | PassedSlowLoadThreshold
    | GotTimeZone Time.Zone


update : Msg -> Model msg -> ( Model msg, Cmd Msg )
update msg model =
    case msg of
        NoOp ->
            ( model, Cmd.none )

        BranchDropdownToggleClicked ->
            let
                state =
                    if model.branchDropdown == BranchDropdown Open then
                        BranchDropdown Closed
                    else
                        BranchDropdown Open
            in
                ( { model | branchDropdown = state }
                , Cmd.none
                )

        BranchDropdownListenClicks ->
            ( { model | branchDropdown = BranchDropdown ListenClicks }
            , Cmd.none
            )

        BranchDropdownClose ->
            ( { model | branchDropdown = BranchDropdown Closed }
            , Cmd.none
            )

        CompletedLoadBuilds (Ok builds) ->
            ( { model | builds = Loaded builds }
            , Cmd.none
            )

        CompletedLoadBuilds (Err error) ->
            ( { model | builds = Failed }
            , Cmd.none
            )

        CompletedLoadTasks (Ok tasks) ->
            ( { model | tasks = Loaded tasks }
            , Cmd.none
            )

        CompletedLoadTasks (Err err) ->
            ( { model | tasks = Failed }
            , Cmd.none
            )

        PassedSlowLoadThreshold ->
            let
                builds =
                    case model.builds of
                        Loading ->
                            LoadingSlowly

                        other ->
                            other

                tasks =
                    case model.tasks of
                        Loading ->
                            LoadingSlowly

                        other ->
                            other
            in
                ( { model
                    | builds = builds
                    , tasks = tasks
                  }
                , Cmd.none
                )

        GotTimeZone tz ->
            ( { model | timeZone = tz }
            , Cmd.none
            )



-- EXPORT


toSession : Model msg -> Session msg
toSession model =
    model.session


toContext : Model msg -> Context msg
toContext model =
    model.context



-- View


view : Model msg -> { title : String, content : Element Msg }
view model =
    let
        device =
            Context.device model.context

        projects =
            Session.projects model.session
    in
        { title = "Project page"
        , content =
            column [ width fill, height fill ]
                [ viewSubHeader device model
                , viewBody model.timeZone device model
                ]
        }


viewErrored : Element msg
viewErrored =
    text "Something went seriously wrong"



-- SubHeader


viewSubHeader : Device -> Model msg -> Element Msg
viewSubHeader device { session, slug, branchDropdown } =
    case Session.projectWithSlug slug session of
        Just project ->
            let
                branches =
                    Session.branches (Project.id project) session
            in
                case device.class of
                    Phone ->
                        viewMobileSubHeader project branches branchDropdown

                    Tablet ->
                        viewDesktopSubHeader project branches branchDropdown

                    Desktop ->
                        viewDesktopSubHeader project branches branchDropdown

                    BigDesktop ->
                        viewDesktopSubHeader project branches branchDropdown

        Nothing ->
            none


viewMobileSubHeader : Project -> List Branch -> BranchDropdown -> Element Msg
viewMobileSubHeader project branches branchDropdown =
    row
        [ width fill
        , height shrink
        , Background.color Palette.white
        , Font.color Palette.white
        , Border.widthEach { top = 1, bottom = 1, left = 0, right = 0 }
        , Border.color Palette.neutral6
        , paddingXY 0 15
        , Border.shadow
            { offset = ( 0, 2 )
            , size = 2
            , blur = 2
            , color = Palette.neutral6
            }
        , below
            (if branchDropdown == BranchDropdown ListenClicks then
                el
                    [ width fill
                    , height (px 9999)
                    , moveDown 1
                    , Background.color Palette.neutral7
                    , Font.color Palette.primary5
                    ]
                    (viewBranchSelectDropdown branches)
             else
                none
            )
        ]
        [ row
            [ width fill
            , centerY
            , Font.color Palette.neutral3
            , spaceEvenly
            ]
            [ el [ width fill, Font.alignLeft, paddingXY 20 0 ] (text <| Project.name project)
            , el [ width fill, paddingXY 20 0 ] (viewBranchSelectButton fill branchDropdown)
            ]
        ]


viewDesktopSubHeader : Project -> List Branch -> BranchDropdown -> Element Msg
viewDesktopSubHeader project branches branchDropdown =
    row
        [ Font.bold
        , Font.size 18
        , width (fill |> maximum 1600)
        , alignRight
        , height (px 65)
        , paddingXY 20 10
        , Font.color Palette.white
        , Background.color Palette.white
        , Font.color Palette.white
        , Border.widthEach { top = 1, bottom = 1, left = 0, right = 0 }
        , Border.color Palette.neutral6
        ]
        [ row
            [ width fill
            , centerY
            , Font.color Palette.neutral3
            ]
            [ el [ width fill ] <|
                el [ alignLeft ] (text <| Project.name project)
            , el
                [ width fill
                , below
                    (if branchDropdown == BranchDropdown ListenClicks then
                        column
                            [ width (fill |> maximum 600 |> minimum 400)
                            , alignRight
                            , Font.size 14
                            , height shrink
                            , Background.color Palette.neutral7
                            , Border.width 1
                            , Border.color Palette.neutral6
                            , Border.rounded 5
                            , moveDown 3
                            , Border.shadow
                                { offset = ( 2, 2 )
                                , size = 1
                                , blur = 2
                                , color = Palette.neutral4
                                }
                            ]
                            [ column
                                [ Border.widthEach { top = 0, left = 0, right = 0, bottom = 1 }
                                , Border.color Palette.neutral6
                                , width fill
                                ]
                                [ row
                                    [ width fill
                                    , spaceEvenly
                                    , paddingEach { bottom = 0, left = 10, right = 10, top = 10 }
                                    ]
                                    [ el
                                        [ alignLeft
                                        , centerY
                                        ]
                                        (text "Switch branches")
                                    , el
                                        [ alignRight
                                        , centerY
                                        , Font.color Palette.neutral4
                                        , pointer
                                        , mouseOver [ Font.color Palette.neutral1 ]
                                        ]
                                        (Icon.x Icon.defaultOptions)
                                    ]
                                , viewBranchSelectDropdown branches
                                ]
                            ]
                     else
                        none
                    )
                ]
                (el
                    [ alignRight
                    ]
                    (viewBranchSelectButton shrink branchDropdown)
                )
            ]
        ]



-- Body


viewBody : Time.Zone -> Device -> Model msg -> Element Msg
viewBody tz device { slug, session, builds, tasks } =
    case Session.projectWithSlug slug session of
        Just project ->
            case device.class of
                Phone ->
                    viewMobileBody tz project tasks builds

                Tablet ->
                    viewDesktopBody tz project tasks builds

                Desktop ->
                    viewDesktopBody tz project tasks builds

                BigDesktop ->
                    viewBigDesktopBody tz project builds

        Nothing ->
            none


viewMobileBody : Time.Zone -> Project -> Status (List Task) -> Status (PaginatedList Build) -> Element Msg
viewMobileBody tz project tasks builds =
    column
        [ width fill
        , height fill
        , Background.color Palette.neutral7
        , centerX
        , spacing 20
        , padding 20
        ]
        [ el [ height shrink, width fill ] (viewProjectDetails project)
        , el [ height shrink, width fill ] (viewRecentTasksContainer tasks)
        , viewProjectBuildsContainer tz builds
        ]


viewDesktopBody : Time.Zone -> Project -> Status (List Task) -> Status (PaginatedList Build) -> Element Msg
viewDesktopBody tz project tasks builds =
    column
        [ width fill
        , height fill
        , Background.color Palette.neutral7
        , paddingXY 20 40
        , spacingXY 20 40
        ]
        [ row
            [ height shrink
            , centerX
            , spacing 20
            , width (fill |> maximum 1600)
            , alignRight
            ]
            [ column
                [ width (fillPortion 4)
                , height fill
                , spacingXY 0 20
                ]
                [ viewProjectDetails project
                ]
            , column
                [ width (fillPortion 6)
                , height fill
                ]
                [ viewRecentTasksContainer tasks ]
            ]
        , row
            [ width fill
            , spacing 20
            ]
            [ viewProjectBuildsContainer tz builds
            ]
        ]


viewBigDesktopBody : Time.Zone -> Project -> Status (PaginatedList Build) -> Element Msg
viewBigDesktopBody tz project builds =
    none



-- Project Branch Tabs


viewProjectBranchTabs : Element msg
viewProjectBranchTabs =
    row [ spaceEvenly, moveUp 55, paddingEach { bottom = 2, left = 0, right = 0, top = 0 } ]
        [ viewProjectBranchTab True "Builds"
        , viewProjectBranchTab False "Commits"
        ]


viewProjectBranchTab : Bool -> String -> Element msg
viewProjectBranchTab isSelected label =
    let
        shadow =
            if isSelected then
                { offset = ( 1, 1 )
                , size = 1
                , blur = 1
                , color = Palette.neutral6
                }
            else
                { offset = ( 0, 0 )
                , size = 0
                , blur = 0
                , color = Palette.transparent
                }
    in
        el
            [ Font.size 14
            , width fill
            , Font.alignLeft
            , paddingXY 25 20
            , Background.color
                (if isSelected then
                    Palette.white
                 else
                    Palette.transparent
                )
            , Border.shadow shadow
            , Border.roundEach { topLeft = 5, topRight = 5, bottomLeft = 0, bottomRight = 0 }
            , Border.color
                (if isSelected then
                    Palette.neutral6
                 else
                    Palette.transparent
                )
            , inFront
                (if isSelected then
                    el
                        [ height (px 10)
                        , alignBottom
                        , moveDown 4
                        , width fill
                        , Background.color Palette.white
                        ]
                        none
                 else
                    none
                )
            ]
            (text label)



-- Project Builds


viewProjectBuildsContainer : Time.Zone -> Status (PaginatedList Build) -> Element Msg
viewProjectBuildsContainer tz buildStatus =
    column
        [ width fill
        , behindContent viewProjectBranchTabs
        , moveDown 50
        ]
        [ row
            [ width fill
            , paddingXY 10 30
            , Font.size 14
            , height (px 300)
            , Background.color Palette.white
            , Border.shadow
                { offset = ( 1, 1 )
                , size = 1
                , blur = 1
                , color = Palette.neutral6
                }
            , Border.roundEach
                { topLeft = 0
                , topRight = 5
                , bottomLeft = 5
                , bottomRight = 5
                }
            ]
            [ viewProjectBuilds tz buildStatus
            ]
        ]


viewProjectBuilds : Time.Zone -> Status (PaginatedList Build) -> Element Msg
viewProjectBuilds tz buildStatus =
    case buildStatus of
        Loaded builds ->
            if PaginatedList.total builds == 0 then
                viewProjectBuildsEmpty
            else
                viewProjectBuildsTable tz builds

        LoadingSlowly ->
            viewLoadingSpinner

        Loading ->
            none

        Failed ->
            viewProjectBuildsFailed


viewProjectBuildsFailed : Element msg
viewProjectBuildsFailed =
    column [ width fill, spacingXY 0 20, Font.color Palette.danger1 ]
        [ row
            [ width shrink
            , centerX
            , Font.size 18
            , Font.color Palette.danger4
            , spacingXY 10 0
            ]
            [ el [] (Icon.alertCircle <| Icon.size 32)
            , paragraph [] [ text "There was a problem loading builds" ]
            ]
        , paragraph
            []
            [ text "If you think this is a bug we'd really appreciate it if you could "
            , newTabLink
                [ Font.color Palette.primary4
                , mouseOver
                    [ Font.color Palette.primary2
                    ]
                ]
                { url = "https://github.com/velocity-ci/velocity", label = text "create an issue on our github" }
            , text " with your configuration files"
            ]
        ]


viewProjectBuildsEmpty : Element msg
viewProjectBuildsEmpty =
    column [ width fill, spacingXY 0 20 ]
        [ row
            [ width shrink
            , centerX
            , Font.size 18
            , Font.color Palette.info4
            , spacingXY 10 0
            ]
            [ el [] (Icon.info <| Icon.size 32)
            , paragraph [] [ text "No builds have been run for branch" ]
            ]
        , paragraph
            [ Font.color Palette.info1 ]
            [ newTabLink
                [ Font.color Palette.primary4
                , mouseOver
                    [ Font.color Palette.primary2
                    ]
                ]
                { url = "https://github.com/velocity-ci/velocity", label = text "Read the documentation" }
            , text " to find out more about builds"
            ]
        ]


viewProjectBuildsTable : Time.Zone -> PaginatedList Build -> Element Msg
viewProjectBuildsTable tz builds =
    indexedTable
        [ width fill
        , height fill
        ]
        { data = PaginatedList.values builds
        , columns =
            [ { header = viewTableHeader (text "Started")
              , width = fillPortion 2
              , view =
                    \i build ->
                        el
                            [ width fill
                            , height (px 60)
                              --                                , Border.widthEach borders
                            , Background.color Palette.neutral7
                            , Border.color Palette.neutral6
                            ]
                        <|
                            el [ centerY, paddingXY 10 0 ] <|
                                (build
                                    |> Build.createdAt
                                    |> Timestamp.format tz
                                    |> text
                                )
              }
              --            , { header = viewTableHeader (text "BaseTask")
              --              , width = fill
              --              , view = \i person -> viewLeftTableCell (text person.task) i
              --              }
              --            , { header = viewTableHeader (text "Commit")
              --              , width = fill
              --              , view =
              --                    \i person ->
              --                        viewLeftTableCell
              --                            (row []
              --                                [ text person.commit
              --                                , text " "
              --                                , row [ Font.heavy ]
              --                                    [ text "("
              --                                    , text person.branch
              --                                    , text ")"
              --                                    ]
              --                                ]
              --                            )
              --                            i
              --              }
              --            , { header = viewTableHeader (text "Status")
              --              , width = fill
              --              , view =
              --                    \i person ->
              --                        viewLeftTableCell
              --                            (case person.status of
              --                                Success ->
              --                                    row [ Font.color Palette.success3, spacingXY 5 0 ]
              --                                        [ el [ Font.heavy ] (Icon.check Icon.defaultOptions)
              --                                        , text "Finished"
              --                                        ]
              --
              --                                Failure ->
              --                                    row [ Font.color Palette.danger3, spacingXY 5 0 ]
              --                                        [ el [ Font.heavy ] (Icon.x Icon.defaultOptions)
              --                                        , text "Finished"
              --                                        ]
              --
              --                                InProgress ->
              --                                    row [] [ text "In progress" ]
              --                            )
              --                            i
              --              }
              --            , { header = viewTableHeader (text "")
              --              , width = shrink
              --              , view =
              --                    \i _ ->
              --                        viewRightTableCell (el [] (Icon.arrowRight Icon.defaultOptions)) i
              --              }
            ]
        }


viewTableHeader : Element msg -> Element msg
viewTableHeader contents =
    el
        [ width fill
        , height (px 40)
        ]
    <|
        el
            [ centerY
            , paddingXY 10 0
            , Font.size 16
            ]
        <|
            contents



-- Project Details


viewProjectDetails : Project -> Element Msg
viewProjectDetails project =
    column
        [ width fill
        , height fill
        , Background.color Palette.white
        , Border.shadow
            { offset = ( 1, 1 )
            , size = 1
            , blur = 1
            , color = Palette.neutral6
            }
        , Border.rounded 5
        ]
        [ el
            [ Font.size 20
            , width fill
            , Font.alignLeft
            , paddingXY 10 20
            , Border.widthEach { top = 0, bottom = 1, left = 0, right = 0 }
            , Border.color Palette.neutral6
            ]
            (text "Details")
        , column [ width fill, centerY ]
            [ row
                [ width fill
                , paddingXY 10 0
                , height (px 50)
                , Border.color Palette.neutral6
                , Font.size 15
                , spaceEvenly
                ]
                [ el [ Font.color Palette.neutral2 ] (text "Repository")
                , row
                    [ spacingXY 5 0
                    , Font.color Palette.primary3
                    , pointer
                    , mouseOver [ Font.color Palette.primary4 ]
                    ]
                    [ el [] (Icon.github Icon.defaultOptions)
                    , el [] (text (Project.name project))
                    ]
                ]
            , row
                [ width fill
                , Border.widthEach { top = 1, bottom = 0, left = 0, right = 0 }
                , paddingXY 10 0
                , Border.color Palette.neutral6
                , Font.size 15
                , spaceEvenly
                , height (px 50)
                , inFront
                    (el
                        [ padding 5
                        , Border.rounded 5
                        , Background.color Palette.neutral6
                        , alignRight
                        , centerY
                        , moveLeft 10
                        , Font.size 13
                        , Font.family [ Font.monospace ]
                        ]
                        (text "master")
                    )
                ]
                [ el [ Font.color Palette.neutral2 ] (text "Default branch")
                ]
            , row
                [ width fill
                , Border.widthEach { top = 1, bottom = 0, left = 0, right = 0 }
                , paddingXY 10 0
                , height (px 50)
                , Border.color Palette.neutral6
                , Font.size 15
                , spaceEvenly
                ]
                [ el [ Font.color Palette.neutral2 ] (text "Updated")
                , el
                    [ Font.color Palette.neutral2
                    ]
                    (text "2 weeks ago")
                ]
            ]
        ]



-- Project Health


viewProjectHealth : Project -> Element Msg
viewProjectHealth project =
    column
        [ width fill
        , height shrink
        , Background.color Palette.white
        , Border.shadow
            { offset = ( 1, 1 )
            , size = 1
            , blur = 1
            , color = Palette.neutral6
            }
        , Border.rounded 5
        ]
        [ el
            [ Font.size 20
            , width fill
            , Font.alignLeft
            , paddingXY 10 20
            ]
            (text "Health")
        , row
            [ width fill
            , Border.widthEach { top = 1, bottom = 0, left = 0, right = 0 }
            , paddingXY 10 0
            , Border.color Palette.neutral6
            , Font.size 15
            , spaceEvenly
            ]
            [ el [ Font.color Palette.neutral2 ] (text "run-unit-tests")
            , el [ Font.color Palette.success3 ] (Icon.sun Icon.defaultOptions)
            ]
        , row
            [ width fill
            , Border.widthEach { top = 1, bottom = 0, left = 0, right = 0 }
            , paddingXY 10 0
            , Border.color Palette.neutral6
            , Font.size 15
            , spaceEvenly
            ]
            [ el [ Font.color Palette.neutral2 ] (text "deploy-master")
            , el [ Font.color Palette.warning3 ] (Icon.cloudRain Icon.defaultOptions)
            ]
        , row
            [ width fill
            , Border.widthEach { top = 1, bottom = 0, left = 0, right = 0 }
            , paddingXY 10 0
            , Border.color Palette.neutral6
            , Font.size 15
            , spaceEvenly
            ]
            [ el [ Font.color Palette.neutral2 ] (text "build-containers")
            , el [ Font.color Palette.danger3 ] (Icon.cloudLightning Icon.defaultOptions)
            ]
        ]


viewRecentTasksContainer : Status (List Task) -> Element Msg
viewRecentTasksContainer tasksStatus =
    column
        [ width fill
        , height fill
        , Background.color Palette.white
        , Border.shadow
            { offset = ( 1, 1 )
            , size = 1
            , blur = 1
            , color = Palette.neutral6
            }
        , Border.rounded 5
        ]
        [ el
            [ Font.size 20
            , width fill
            , Font.alignLeft
            , paddingXY 10 20
            ]
            (text "Tasks")
        , row
            [ width fill
            , height fill
            , Border.widthEach { top = 1, left = 0, right = 0, bottom = 0 }
            , Border.color Palette.neutral6
            , paddingXY 10 10
            , spacing 10
            ]
            (viewRecentTasks tasksStatus)
        ]


viewRecentTasks : Status (List Task) -> List (Element Msg)
viewRecentTasks tasksStatus =
    case tasksStatus of
        Loaded tasks ->
            List.map viewRecentTaskItem (List.take 1 tasks)

        Loading ->
            [ none ]

        LoadingSlowly ->
            [ viewLoadingSpinner ]

        Failed ->
            [ viewRecentTasksFailed ]


viewRecentTasksFailed : Element msg
viewRecentTasksFailed =
    column [ width fill, spacingXY 0 20, padding 20, height fill, Font.color Palette.danger1 ]
        [ row
            [ width shrink
            , centerX
            , Font.size 18
            , Font.color Palette.danger4
            , spacingXY 10 0
            ]
            [ el [] (Icon.alertCircle <| Icon.size 32)
            , text "There was a problem loading tasks"
            ]
        , textColumn [ width fill, spacingXY 0 20, Font.size 14 ]
            [ paragraph
                [ width fill
                ]
                [ text "Velocity failed loading tasks from the commit at "
                , el [ Font.heavy ] (text "HEAD")
                , text " of the branch."
                ]
            , paragraph []
                [ text "If you think this is a bug we'd really appreciate it if you could "
                , newTabLink
                    [ Font.color Palette.primary4
                    , mouseOver
                        [ Font.color Palette.primary2
                        ]
                    ]
                    { url = "https://github.com/velocity-ci/velocity", label = text "create an issue on our github" }
                , text " with your configuration files and any other information that you think would be helpful"
                ]
            ]
        ]


viewRecentTaskItem : Task -> Element msg
viewRecentTaskItem task =
    let
        taskName =
            task
                |> Task.name
                |> TaskName.toString
    in
        column
            [ width fill
            , paddingXY 10 20
            , Font.size 15
            , spacingXY 0 10
            , height fill
            , centerY
            , pointer
            , mouseOver
                [ Background.color Palette.neutral7
                ]
            , Border.width 1
            , Border.color Palette.neutral6
            , Border.rounded 10
            , clipX
            ]
            [ el [ height fill, Font.color Palette.danger4, width shrink, centerX ] (Icon.xCircle <| Icon.size 38)
            , paragraph [ width fill ] [ el [ Font.color Palette.neutral1, width fill, centerX, Font.alignLeft ] (text taskName) ]
            , el [ height fill, Font.color Palette.neutral3, width shrink, centerX ] (el [ alignBottom ] (text "3 months ago"))
            ]



-- Task Container


viewTasksContainer : Model msg -> Element Msg
viewTasksContainer model =
    column
        [ width fill
        , height shrink
        , padding 10
        ]
        [ viewTasksList
        ]


viewBranchSelectButton : Length -> BranchDropdown -> Element Msg
viewBranchSelectButton widthLength (BranchDropdown state) =
    el
        [ width fill
        , height fill
        ]
        (Button.button BranchDropdownToggleClicked
            { leftIcon = Nothing
            , rightIcon = Nothing
            , centerLeftIcon = Nothing
            , centerRightIcon = Nothing
            , scheme =
                if state == Closed then
                    Button.Secondary
                else
                    Button.Primary
            , content =
                row [ spacingXY 5 0 ]
                    [ el [] (text "Branch:")
                    , el [ Font.heavy ] (text "master")
                    , el [] (Icon.arrowDown { size = 12, sizeUnit = Icon.Pixels, strokeWidth = 1 })
                    ]
            , size = Button.Small
            , widthLength = widthLength
            , heightLength = shrink
            , disabled = False
            }
        )


viewBranchSelectDropdown : List Branch -> Element Msg
viewBranchSelectDropdown branches =
    column [ width fill ]
        [ row
            [ width fill
            , padding 10
            ]
            [ el
                [ Background.color Palette.white
                , width fill
                ]
                (Input.search
                    { leftIcon = Just Icon.search
                    , rightIcon = Nothing
                    , label = Input.labelHidden "Search for a branch"
                    , placeholder = Just "Find a branch..."
                    , dirty = False
                    , value = ""
                    , problems = []
                    , onChange = always NoOp
                    }
                )
            ]
        , column
            [ width fill
            , paddingEach { top = 0, left = 0, right = 0, bottom = 0 }
            ]
            (List.indexedMap
                (\i b ->
                    let
                        rounded =
                            if i + 1 == List.length branches then
                                { topLeft = 0
                                , topRight = 0
                                , bottomLeft = 5
                                , bottomRight = 5
                                }
                            else
                                { topLeft = 0
                                , topRight = 0
                                , bottomLeft = 0
                                , bottomRight = 0
                                }
                    in
                        viewBranchSelectDropdownItem rounded b
                )
                branches
            )
        ]


viewBranchSelectDropdownItem :
    { topLeft : Int
    , topRight : Int
    , bottomLeft : Int
    , bottomRight : Int
    }
    -> Branch
    -> Element Msg
viewBranchSelectDropdownItem rounded branch =
    row
        [ Border.widthEach { top = 1, bottom = 0, right = 0, left = 0 }
        , Border.color Palette.neutral6
        , width fill
        , padding 10
        , Background.color Palette.white
        , spacingXY 10 0
        , Font.color Palette.neutral1
        , Border.roundEach rounded
        , pointer
        , mouseOver
            [ Background.color Palette.primary4
            , Font.color Palette.neutral7
            ]
        ]
        [ el [ width (px 16), centerY ] none
        , el [ width fill, centerY, Font.alignLeft, clipX ] (Branch.text branch)
        ]


viewTasksList : Element msg
viewTasksList =
    column [ width fill, spacingXY 0 10, paddingXY 0 10 ]
        []


viewTask : Element msg
viewTask =
    row
        [ width fill
        , padding 15
        , Background.color Palette.neutral7
        ]
        [ text "run-unit-tests" ]



-- Timeline


avatarIcon : Element msg
avatarIcon =
    el
        [ width (px 30)
        , height (px 30)
        , Border.rounded 180
        , Background.image <| Asset.src Asset.defaultAvatar
        ]
        none


viewTimeline : Element msg
viewTimeline =
    column
        [ width fill
        , height shrink
        , Background.color Palette.white
        , Border.shadow
            { offset = ( 1, 1 )
            , size = 1
            , blur = 1
            , color = Palette.neutral6
            }
        , Border.rounded 5
        , Font.size 14
        , Font.alignLeft
        ]
        [ el
            [ Font.size 20
            , width fill
            , Font.alignLeft
            , paddingXY 10 20
            , Border.widthEach { bottom = 1, right = 0, top = 0, left = 0 }
            , Border.color Palette.neutral6
            ]
            (text "Timeline")
        , column
            [ width fill
            , behindContent
                (el
                    [ height fill
                    , width (px 1)
                    , Background.color Palette.neutral6
                    , moveRight 24
                    ]
                    none
                )
            ]
            [ viewCommitListRowOne { top = 0, bottom = 1 }
            , viewCommitListRowTwo { top = 0, bottom = 1 }
            , viewCommitListRowThree { top = 0, bottom = 0 }
            ]
        ]


viewTimelineRow : { topBorder : Int, bottomBorder : Int } -> Element msg -> Element msg
viewTimelineRow { topBorder, bottomBorder } content =
    content


viewCommitListRowOne : { top : Int, bottom : Int } -> Element msg
viewCommitListRowOne { top, bottom } =
    row
        [ spacingXY 15 0
        , width fill
        , Border.widthEach { bottom = bottom, top = top, left = 0, right = 0 }
        , paddingXY 10 17
        , Border.color Palette.neutral6
        , height (px 50)
        ]
        [ el
            [ width shrink
            , height shrink
            ]
            avatarIcon
        , wrappedRow [ width fill, height shrink, spacing 5 ]
            [ el [ Font.color Palette.neutral1, Font.heavy ] (text "VJ")
            , el [ Font.color Palette.neutral3 ] (text "pushed")
            , el [ Font.color Palette.neutral1, Font.heavy ] (text "WIP - Frontend compiling")
            , el [ Font.color Palette.neutral3 ] (text "to")
            , el [ Font.color Palette.neutral1, Font.heavy ] (text "CP-179-multicargo")
            , el [ Font.size 14, Font.color Palette.neutral5, alignRight ] (text "2 days ago")
            ]
        ]


viewCommitListRowTwo : { top : Int, bottom : Int } -> Element msg
viewCommitListRowTwo { top, bottom } =
    row
        [ spacingXY 15 0
        , width fill
        , Border.widthEach { bottom = bottom, top = top, left = 0, right = 0 }
        , paddingXY 10 17
        , Border.color Palette.neutral6
        , height (px 50)
        ]
        [ el
            [ width shrink
            , height shrink
            ]
            avatarIcon
        , wrappedRow [ width fill, height shrink, spacing 5 ]
            [ el
                [ padding 5
                , Background.color Palette.success5
                , Font.color Palette.white
                , Font.heavy
                , Font.size 16
                , Font.variant Font.smallCaps
                , Border.rounded 5
                , width (px 60)
                ]
                (el
                    [ centerX
                    , centerY
                    ]
                    (text "success")
                )
            , el [ Font.color Palette.neutral1, Font.heavy ] (text "Eddy")
            , el [ Font.color Palette.neutral3 ] (text "ran")
            , el [ Font.color Palette.neutral1, Font.heavy ] (text "run-unit-tests")
            , el [ Font.color Palette.neutral3 ] (text "on")
            , el [ Font.color Palette.neutral1, Font.heavy ] (text "master")
            , el [ Font.size 14, Font.color Palette.neutral5, alignRight ] (text "2 days ago")
            ]
        ]


viewCommitListRowThree : { top : Int, bottom : Int } -> Element msg
viewCommitListRowThree { top, bottom } =
    row
        [ spacingXY 15 0
        , width fill
        , Border.widthEach { bottom = bottom, top = top, left = 0, right = 0 }
        , paddingXY 10 17
        , Border.color Palette.neutral6
        , height (px 50)
        ]
        [ el
            [ width shrink
            , height shrink
            ]
            avatarIcon
        , wrappedRow [ width fill, height shrink, spacing 5 ]
            [ el
                [ padding 5
                , Background.color Palette.danger5
                , Font.color Palette.white
                , Font.heavy
                , Font.size 16
                , Font.variant Font.smallCaps
                , Border.rounded 5
                , width (px 60)
                ]
                (el
                    [ centerX
                    , centerY
                    ]
                    (text "failure")
                )
            , el [ Font.color Palette.neutral1, Font.heavy ] (text "Eddy")
            , el [ Font.color Palette.neutral3 ] (text "ran")
            , el [ Font.color Palette.neutral1, Font.heavy ] (text "run-unit-tests")
            , el [ Font.color Palette.neutral3 ] (text "on")
            , el [ Font.color Palette.neutral1, Font.heavy ] (text "master")
            , el [ Font.size 14, Font.color Palette.neutral5, alignRight ] (text "2 days ago")
            ]
        ]



-- MISC


viewLoadingSpinner : Element msg
viewLoadingSpinner =
    el
        [ centerX
        , centerY
        , Font.color Palette.primary4
        ]
    <|
        Loading.icon { width = 40, height = 40 }
