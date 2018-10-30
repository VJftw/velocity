module Data.BuildStep exposing (BuildStep, Id(..), Status(..), decoder, idParser, idToString, statusDecoder)

import Data.BuildStream as BuildStream exposing (BuildStream)
import Data.Commit as Commit
import Data.Helpers exposing (stringToDateTime)
import Data.Project as Project
import Data.Task as Task
import Json.Decode as Decode exposing (Decoder, int, string)
import Json.Decode.Pipeline as Pipeline exposing (custom, decode, hardcoded, optional, required)
import Time.DateTime as DateTime exposing (DateTime)
import UrlParser


type alias BuildStep =
    { id : Id
    , status : Status
    , number : Int
    , streams : List BuildStream
    , startedAt : Maybe DateTime
    , updatedAt : Maybe DateTime
    }



-- SERIALIZATION --


decoder : Decoder BuildStep
decoder =
    decode BuildStep
        |> required "id" (Decode.map Id string)
        |> required "status" statusDecoder
        |> required "number" Decode.int
        |> required "streams" (Decode.list BuildStream.decoder)
        |> required "startedAt" (Decode.maybe stringToDateTime)
        |> required "updatedAt" (Decode.maybe stringToDateTime)


statusDecoder : Decoder Status
statusDecoder =
    Decode.string
        |> Decode.andThen
            (\status ->
                case status of
                    "waiting" ->
                        Decode.succeed Waiting

                    "failed" ->
                        Decode.succeed Failed

                    "running" ->
                        Decode.succeed Running

                    "success" ->
                        Decode.succeed Success

                    unknown ->
                        Decode.fail <| "Unknown status: " ++ unknown
            )



-- IDENTIFIERS --


idParser : UrlParser.Parser (Id -> a) a
idParser =
    UrlParser.custom "ID" (Ok << Id)


type Status
    = Waiting
    | Failed
    | Running
    | Success


type Id
    = Id String


idToString : Id -> String
idToString (Id id) =
    id
