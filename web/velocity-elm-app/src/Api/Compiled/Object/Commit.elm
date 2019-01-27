-- Do not manually edit this file, it was auto-generated by dillonkearns/elm-graphql
-- https://github.com/dillonkearns/elm-graphql


module Api.Compiled.Object.Commit exposing (author, branches, gpgFingerprint, id, message, sha, tasks)

import Api.Compiled.InputObject
import Api.Compiled.Interface
import Api.Compiled.Object
import Api.Compiled.Scalar
import Api.Compiled.Union
import Graphql.Internal.Builder.Argument as Argument exposing (Argument)
import Graphql.Internal.Builder.Object as Object
import Graphql.Internal.Encode as Encode exposing (Value)
import Graphql.Operation exposing (RootMutation, RootQuery, RootSubscription)
import Graphql.OptionalArgument exposing (OptionalArgument(..))
import Graphql.SelectionSet exposing (SelectionSet)
import Json.Decode as Decode


author : SelectionSet decodesTo Api.Compiled.Object.CommitAuthor -> SelectionSet decodesTo Api.Compiled.Object.Commit
author object_ =
    Object.selectionForCompositeField "author" [] object_ identity


branches : SelectionSet decodesTo Api.Compiled.Object.Branch -> SelectionSet (List decodesTo) Api.Compiled.Object.Commit
branches object_ =
    Object.selectionForCompositeField "branches" [] object_ (identity >> Decode.list)


gpgFingerprint : SelectionSet String Api.Compiled.Object.Commit
gpgFingerprint =
    Object.selectionForField "String" "gpgFingerprint" [] Decode.string


{-| The ID of an object
-}
id : SelectionSet Api.Compiled.Scalar.Id Api.Compiled.Object.Commit
id =
    Object.selectionForField "Scalar.Id" "id" [] (Object.scalarDecoder |> Decode.map Api.Compiled.Scalar.Id)


message : SelectionSet String Api.Compiled.Object.Commit
message =
    Object.selectionForField "String" "message" [] Decode.string


sha : SelectionSet String Api.Compiled.Object.Commit
sha =
    Object.selectionForField "String" "sha" [] Decode.string


tasks : SelectionSet decodesTo Api.Compiled.Object.Task -> SelectionSet (List decodesTo) Api.Compiled.Object.Commit
tasks object_ =
    Object.selectionForCompositeField "tasks" [] object_ (identity >> Decode.list)
