import Websocket from "./Websocket/websocket.js";
import TrainStore from "./stores/trainStore.js";

const app = {
    websocket: new Websocket(`ws://${window.location.hostname}:${window.location.port}/ws`),
    trainStore: null
}

window.addEventListener("load", function(_) {
    app.websocket.init().then(() => {
        app.trainStore = new TrainStore((trainName, trainFields) => {
            const message = {
                train: trainName,
                values: JSON.stringify(trainFields)
            }

            const messageJson = JSON.stringify(message)
            app.websocket.send("patchTrain", messageJson)
        },)

        app.websocket.register("getTrains", (trainsJson) => {
            const trains = JSON.parse(trainsJson)
            app.trainStore.set(trains)
        })
        app.websocket.register("updateTrain", (trainJson) => {
            const train = JSON.parse(trainJson)
            app.trainStore.addOrUpdate(train)
        })
        app.websocket.register("removeTrain", (trainName) => {
            app.trainStore.remove(trainName)
        })

        // load all currently known trains
        app.websocket.send("getTrains")
    })
})