module Main exposing (Document, Model, Msg(..), Rect, decodeRect, encodeRect, init, main, subscriptions, update, view)

import Browser
import Html exposing (Attribute, Html, div, input, text)
import Html.Attributes exposing (..)
import Html.Events exposing (on, onInput)
import Json.Decode as Decode
import Json.Encode as Encode
import Maybe


image =
    "/static/understanding_the_apple_ii_0007.png"



-- MAIN


main =
    Browser.document
        { init = init
        , update = update
        , subscriptions = subscriptions
        , view = view
        }



-- MODEL


type alias Model =
    { content : String
    , rects : List Rect
    }


init : () -> ( Model, Cmd Msg )
init _ =
    ( { content = "", rects = initialRects }
    , Cmd.none
    )



-- UPDATE


type Msg
    = ChangeString String
    | ChangeRects (List Rect)


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        ChangeString newContent ->
            ( { model | content = newContent }, Cmd.none )

        ChangeRects newRects ->
            ( { model | rects = newRects }, Cmd.none )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


type alias Document msg =
    { title : String
    , body : List (Html msg)
    }


view : Model -> Document Msg
view model =
    { title = "Rectangle Editor v1"
    , body =
        [ Html.node "rekt-editor"
            [ attribute "bgimage" image
            , attribute "bgimagewidth" "850"
            , attribute "bgimageheight" "1100"
            , attribute "bgrotate" "0.8"
            , attribute "bgx" "69"
            , attribute "bgy" "75"
            , attribute "bgwidth" "363"
            , attribute "bgheight" "981"
            , property "rectangles" (encodeRects model.rects)
            , on "rectChange" rectChangeDecoder
            ]
            []
        , Html.span [ attribute "id" "rects" ]
            [ Html.pre [] [ text (Encode.encode 4 (encodeRects model.rects)) ] ]
        ]
    }



-- RECTS


rectChangeDecoder : Decode.Decoder Msg
rectChangeDecoder =
    Decode.map ChangeRects customEventDecoder


customEventDecoder : Decode.Decoder (List Rect)
customEventDecoder =
    Decode.field "detail" decodeRects


type alias Rect =
    { id : String
    , x : Float
    , y : Float
    , width : Float
    , height : Float
    , rotate : Float
    , active : Bool
    }


initialRects : List Rect
initialRects =
    [ { id = "rect-5ca1ab1e"
      , width = 30
      , height = 15
      , x = 142
      , y = 271
      , rotate = 20
      , active = False
      }
    , { id = "rect-f01dab1e"
      , width = 363.003
      , height = 127.546
      , x = 68.0957
      , y = 76.2188
      , rotate = -0.847341
      , active = True
      }
    ]


encodeRect : Rect -> Encode.Value
encodeRect rect =
    Encode.object
        [ ( "id", Encode.string rect.id )
        , ( "x", Encode.float rect.x )
        , ( "y", Encode.float rect.y )
        , ( "width", Encode.float rect.width )
        , ( "height", Encode.float rect.height )
        , ( "rotate", Encode.float rect.rotate )
        , ( "active", Encode.bool rect.active )
        ]


encodeRects : List Rect -> Encode.Value
encodeRects =
    Encode.list encodeRect


decodeRect : Decode.Decoder Rect
decodeRect =
    Decode.map7 Rect
        (Decode.field "id" Decode.string)
        (Decode.field "x" Decode.float)
        (Decode.field "y" Decode.float)
        (Decode.field "width" Decode.float)
        (Decode.field "height" Decode.float)
        (Decode.field "rotate" Decode.float)
        (Decode.map (Maybe.withDefault False) (Decode.maybe (Decode.field "active" Decode.bool)))


decodeRects : Decode.Decoder (List Rect)
decodeRects =
    Decode.list decodeRect
