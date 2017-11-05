module Request.Commit exposing (list, get, tasks, task, builds, createBuild)

import Data.AuthToken as AuthToken exposing (AuthToken, withAuthorization)
import Data.Project as Project exposing (Project)
import Data.Commit as Commit exposing (Commit)
import Data.Task as Task exposing (Task)
import Data.Branch as Branch exposing (Branch)
import Data.Build as Build exposing (Build)
import Data.CommitResults as CommitResults
import Json.Decode as Decode
import Json.Encode as Encode
import Request.Helpers exposing (apiUrl)
import HttpBuilder exposing (RequestBuilder, withBody, withExpect, withQueryParams)
import Util exposing ((=>))
import Http


baseUrl : String
baseUrl =
    "/projects"



-- COMMITS --


list :
    Project.Id
    -> Maybe Branch
    -> Int
    -> Int
    -> Maybe AuthToken
    -> Http.Request CommitResults.Results
list id maybeBranch amount page maybeToken =
    let
        expect =
            CommitResults.decoder
                |> Http.expectJson

        branchParam queryParams =
            case maybeBranch of
                Just (Branch.Name branch) ->
                    ( "branch", branch ) :: queryParams

                _ ->
                    queryParams

        amountParam queryParams =
            ( "amount", toString amount ) :: queryParams

        pageParam queryParams =
            ( "page", toString page ) :: queryParams

        queryParams =
            []
                |> branchParam
                |> amountParam
                |> pageParam
    in
        apiUrl (baseUrl ++ "/" ++ Project.idToString id ++ "/commits")
            |> HttpBuilder.get
            |> HttpBuilder.withExpect expect
            |> HttpBuilder.withQueryParams queryParams
            |> withAuthorization maybeToken
            |> HttpBuilder.toRequest



-- GET --


get : Project.Id -> Commit.Hash -> Maybe AuthToken -> Http.Request Commit
get id hash maybeToken =
    let
        expect =
            Commit.decoder
                |> Http.expectJson

        urlPieces =
            [ baseUrl
            , Project.idToString id
            , "commits"
            , Commit.hashToString hash
            ]
    in
        apiUrl (String.join "/" urlPieces)
            |> HttpBuilder.get
            |> HttpBuilder.withExpect expect
            |> withAuthorization maybeToken
            |> HttpBuilder.toRequest



-- TASKS --


tasks : Project.Id -> Commit.Hash -> Maybe AuthToken -> Http.Request (List Task)
tasks id hash maybeToken =
    let
        expect =
            Task.decoder
                |> Decode.list
                |> Http.expectJson

        urlPieces =
            [ baseUrl
            , Project.idToString id
            , "commits"
            , Commit.hashToString hash
            , "tasks"
            ]
    in
        apiUrl (String.join "/" urlPieces)
            |> HttpBuilder.get
            |> HttpBuilder.withExpect expect
            |> withAuthorization maybeToken
            |> HttpBuilder.toRequest


task : Project.Id -> Commit.Hash -> Task.Name -> Maybe AuthToken -> Http.Request Task
task id hash name maybeToken =
    let
        expect =
            Task.decoder
                |> Http.expectJson

        urlPieces =
            [ baseUrl
            , Project.idToString id
            , "commits"
            , Commit.hashToString hash
            , "tasks"
            , Task.nameToString name
            ]
    in
        apiUrl (String.join "/" urlPieces)
            |> HttpBuilder.get
            |> HttpBuilder.withExpect expect
            |> withAuthorization maybeToken
            |> HttpBuilder.toRequest



-- BUILDS --


builds : Project.Id -> Commit.Hash -> Maybe AuthToken -> Http.Request (List Build)
builds id hash maybeToken =
    let
        expect =
            Build.decoder
                |> Decode.list
                |> Http.expectJson

        urlPieces =
            [ baseUrl
            , Project.idToString id
            , "commits"
            , Commit.hashToString hash
            , "builds"
            ]
    in
        apiUrl (String.join "/" urlPieces)
            |> HttpBuilder.get
            |> HttpBuilder.withExpect expect
            |> withAuthorization maybeToken
            |> HttpBuilder.toRequest


createBuild : Project.Id -> Commit.Hash -> Task.Name -> List ( String, String ) -> AuthToken -> Http.Request Build
createBuild id hash taskName params token =
    let
        expect =
            Build.decoder
                |> Http.expectJson

        urlPieces =
            [ baseUrl
            , Project.idToString id
            , "commits"
            , Commit.hashToString hash
            , "builds"
            ]

        encodedParams =
            let
                enc =
                    params
                        |> List.map (\( field, value ) -> field => Encode.string value)
            in
                if (List.length params) > 0 then
                    Encode.object enc
                else
                    Encode.list []

        encodedBody =
            Encode.object
                [ "taskName" => (Task.nameToString taskName |> Encode.string)
                , "params" => encodedParams
                ]

        body =
            encodedBody
                |> Http.jsonBody
    in
        apiUrl (String.join "/" urlPieces)
            |> HttpBuilder.post
            |> HttpBuilder.withExpect expect
            |> withAuthorization (Just token)
            |> withBody body
            |> HttpBuilder.toRequest
