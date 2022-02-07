package main

import (
    "fmt"
    "github.com/sebastiansalte/elv/main.go"
)
func main(){

    fmt.Println(state.ViewState())

    //HER KOMMER KYLLINGEN
    err := state.PutInBoat("HS kylling")
    if err != nil {
    fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())

    err = state.CrossRiverTo("east")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())



    //DETTE ER KANSKJE EN BÅT
    err = state.PutInBoat("HS")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())

    err = state.CrossRiverTo("west")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())



    //EN SEKK MED KORN??
    err = state.PutInBoat("HS korn")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())

    err = state.CrossRiverTo("east")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())



    //NOK EN KYLLING
    err = state.PutInBoat("HS kylling")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())

    err = state.CrossRiverTo("west")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())



    //OG EN REV... GAMMELT NAVN FOR JAY
    err = state.PutInBoat("HS rev")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())

    err = state.CrossRiverTo("east")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())



    //I EN BÅT
    err = state.PutInBoat("HS")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())

    err = state.CrossRiverTo("west")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())



    //DEN TREDJE KYLLING
    err = state.PutInBoat("HS kylling")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())

    err = state.CrossRiverTo("east")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(state.ViewState())


    fmt.Println("\nFerdig")

}