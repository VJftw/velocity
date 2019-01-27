-- Do not manually edit this file, it was auto-generated by dillonkearns/elm-graphql
-- https://github.com/dillonkearns/elm-graphql


module Api.Compiled.Object.Branch exposing (CommitsOptionalArguments, commitAmount, commits, id, name, tasks)

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


commitAmount : SelectionSet Int Api.Compiled.Object.Branch
commitAmount =
    Object.selectionForField "Int" "commitAmount" [] Decode.int


type alias CommitsOptionalArguments =
    { after : OptionalArgument String
    , before : OptionalArgument String
    , first : OptionalArgument Int
    , last : OptionalArgument Int
    }


commits : (CommitsOptionalArguments -> CommitsOptionalArguments) -> SelectionSet decodesTo Api.Compiled.Object.CommitConnection -> SelectionSet (Maybe decodesTo) Api.Compiled.Object.Branch
commits fillInOptionals object_ =
    let
        filledInOptionals =
            fillInOptionals { after = Absent, before = Absent, first = Absent, last = Absent }

        optionalArgs =
            [ Argument.optional "after" filledInOptionals.after Encode.string, Argument.optional "before" filledInOptionals.before Encode.string, Argument.optional "first" filledInOptionals.first Encode.int, Argument.optional "last" filledInOptionals.last Encode.int ]
                |> List.filterMap identity
    in
    Object.selectionForCompositeField "commits" optionalArgs object_ (identity >> Decode.nullable)


{-| The ID of an object
-}
id : SelectionSet Api.Compiled.Scalar.Id Api.Compiled.Object.Branch
id =
    Object.selectionForField "Scalar.Id" "id" [] (Object.scalarDecoder |> Decode.map Api.Compiled.Scalar.Id)


name : SelectionSet String Api.Compiled.Object.Branch
name =
    Object.selectionForField "String" "name" [] Decode.string


tasks : SelectionSet decodesTo Api.Compiled.Object.Task -> SelectionSet (List decodesTo) Api.Compiled.Object.Branch
tasks object_ =
    Object.selectionForCompositeField "tasks" [] object_ (identity >> Decode.list)
