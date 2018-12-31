-- Do not manually edit this file, it was auto-generated by dillonkearns/elm-graphql
-- https://github.com/dillonkearns/elm-graphql


module Api.Compiled.Mutation exposing (SignInRequiredArguments, SignUpOptionalArguments, signIn, signUp)

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
import Json.Decode as Decode exposing (Decoder)


type alias SignInRequiredArguments =
    { password : String
    , username : String
    }


{-| Sign in
-}
signIn : SignInRequiredArguments -> SelectionSet decodesTo Api.Compiled.Object.SessionPayload -> SelectionSet (Maybe decodesTo) RootMutation
signIn requiredArgs object_ =
    Object.selectionForCompositeField "signIn" [ Argument.required "password" requiredArgs.password Encode.string, Argument.required "username" requiredArgs.username Encode.string ] object_ (identity >> Decode.nullable)


type alias SignUpOptionalArguments =
    { password : OptionalArgument String
    , username : OptionalArgument String
    }


{-| Sign up
-}
signUp : (SignUpOptionalArguments -> SignUpOptionalArguments) -> SelectionSet decodesTo Api.Compiled.Object.UserPayload -> SelectionSet (Maybe decodesTo) RootMutation
signUp fillInOptionals object_ =
    let
        filledInOptionals =
            fillInOptionals { password = Absent, username = Absent }

        optionalArgs =
            [ Argument.optional "password" filledInOptionals.password Encode.string, Argument.optional "username" filledInOptionals.username Encode.string ]
                |> List.filterMap identity
    in
        Object.selectionForCompositeField "signUp" optionalArgs object_ (identity >> Decode.nullable)
