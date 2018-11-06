module Session exposing (InitError, Session, changes, cred, errorToString, fromViewer, navKey, projects, viewer)

import Api exposing (BaseUrl, Cred)
import Browser.Navigation as Nav
import Http
import Project exposing (Project)
import Task exposing (Task)
import Viewer exposing (Viewer)



-- TYPES


type Session
    = LoggedIn Nav.Key Viewer (List Project)
    | Guest Nav.Key


type InitError
    = HttpError Http.Error



-- INFO


viewer : Session -> Maybe Viewer
viewer session =
    case session of
        LoggedIn _ val _ ->
            Just val

        Guest _ ->
            Nothing


projects : Session -> List Project
projects session =
    case session of
        LoggedIn _ _ projects_ ->
            projects_

        Guest _ ->
            []


cred : Session -> Maybe Cred
cred session =
    case session of
        LoggedIn _ val _ ->
            Just (Viewer.cred val)

        Guest _ ->
            Nothing


navKey : Session -> Nav.Key
navKey session =
    case session of
        LoggedIn key _ _ ->
            key

        Guest key ->
            key


errorToString : InitError -> String
errorToString (HttpError httpError) =
    case httpError of
        Http.BadUrl error ->
            "Bad URL: " ++ error

        Http.NetworkError ->
            "Network Error"

        Http.BadStatus _ ->
            "Bad Status"

        Http.BadPayload payload _ ->
            "Bad Payload: " ++ payload

        Http.Timeout ->
            "Timeout"



-- CHANGES


changes : (Task InitError Session -> msg) -> BaseUrl -> Session -> Sub msg
changes toMsg baseUrl session =
    Api.viewerChanges (fromViewer (navKey session) baseUrl >> toMsg) Viewer.decoder


fromViewer : Nav.Key -> BaseUrl -> Maybe Viewer -> Task InitError Session
fromViewer key baseUrl maybeViewer =
    -- It's stored in localStorage as a JSON String;
    -- first decode the Value as a String, then
    -- decode that String as JSON.
    -- If the person is logged in we will attempt to get a
    case maybeViewer of
        Just viewerVal ->
            let
                credVal =
                    Viewer.cred viewerVal
            in
            Project.list (Just credVal) baseUrl
                |> Http.toTask
                |> Task.mapError HttpError
                |> Task.map (LoggedIn key viewerVal)

        Nothing ->
            Task.succeed (Guest key)
