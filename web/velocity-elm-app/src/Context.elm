module Context exposing (Context, baseUrl, device, start, windowResize)

{-| The runtime context of the application.
-}

import Api exposing (BaseUrl)
import Element exposing (Device)
import Email exposing (Email)
import Http
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Pipeline exposing (custom, required)
import Json.Encode as Encode exposing (Value)
import Project exposing (Project)
import Task exposing (Task)
import Username exposing (Username)



-- TYPES


type Context
    = Context BaseUrl Device


start : BaseUrl -> { width : Int, height : Int } -> Context
start baseUrl_ dimensions =
    let
        device_ =
            Element.classifyDevice dimensions
    in
    Context baseUrl_ device_



-- INFO


baseUrl : Context -> BaseUrl
baseUrl (Context val _) =
    val


device : Context -> Device
device (Context _ val) =
    val



-- CHANGES


windowResize : { width : Int, height : Int } -> Context -> Context
windowResize dimensions (Context baseUrl_ _) =
    Context baseUrl_ (Element.classifyDevice dimensions)
