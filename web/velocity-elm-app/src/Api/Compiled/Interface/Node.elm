-- Do not manually edit this file, it was auto-generated by dillonkearns/elm-graphql
-- https://github.com/dillonkearns/elm-graphql


module Api.Compiled.Interface.Node exposing (Fragments, fragments, id, maybeFragments)

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
import Graphql.SelectionSet exposing (FragmentSelectionSet(..), SelectionSet(..))
import Json.Decode as Decode


type alias Fragments decodesTo =
    { onProject : SelectionSet decodesTo Api.Compiled.Object.Project
    , onBranch : SelectionSet decodesTo Api.Compiled.Object.Branch
    , onCommit : SelectionSet decodesTo Api.Compiled.Object.Commit
    , onCommitAuthor : SelectionSet decodesTo Api.Compiled.Object.CommitAuthor
    , onTask : SelectionSet decodesTo Api.Compiled.Object.Task
    }


{-| Build an exhaustive selection of type-specific fragments.
-}
fragments :
    Fragments decodesTo
    -> SelectionSet decodesTo Api.Compiled.Interface.Node
fragments selections =
    Object.exhuastiveFragmentSelection
        [ Object.buildFragment "Project" selections.onProject
        , Object.buildFragment "Branch" selections.onBranch
        , Object.buildFragment "Commit" selections.onCommit
        , Object.buildFragment "CommitAuthor" selections.onCommitAuthor
        , Object.buildFragment "Task" selections.onTask
        ]


{-| Can be used to create a non-exhuastive set of fragments by using the record
update syntax to add `SelectionSet`s for the types you want to handle.
-}
maybeFragments : Fragments (Maybe decodesTo)
maybeFragments =
    { onProject = Graphql.SelectionSet.empty |> Graphql.SelectionSet.map (\_ -> Nothing)
    , onBranch = Graphql.SelectionSet.empty |> Graphql.SelectionSet.map (\_ -> Nothing)
    , onCommit = Graphql.SelectionSet.empty |> Graphql.SelectionSet.map (\_ -> Nothing)
    , onCommitAuthor = Graphql.SelectionSet.empty |> Graphql.SelectionSet.map (\_ -> Nothing)
    , onTask = Graphql.SelectionSet.empty |> Graphql.SelectionSet.map (\_ -> Nothing)
    }


{-| The id of the object.
-}
id : SelectionSet Api.Compiled.Scalar.Id Api.Compiled.Interface.Node
id =
    Object.selectionForField "Scalar.Id" "id" [] (Object.scalarDecoder |> Decode.map Api.Compiled.Scalar.Id)
