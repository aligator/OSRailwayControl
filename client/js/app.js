import Websocket from "./Websocket/websocket.js";
import TrainStore from "./stores/trainStore.js";

const app = {
    websocket: new Websocket(`ws://${window.location.hostname}:${window.location.port}/wss`),
    trainStore: null
}

window.addEventListener("load", function(_) {
    app.websocket.init().then(() => {
        app.trainStore = new TrainStore((train, speed) => {
            const message = {
                train,
                speed
            }

            const messageJson = JSON.stringify(message)

            app.websocket.send("setSpeed", messageJson)
        })

        app.websocket.register("getTrains", (trainsJson) => {
            const trains = JSON.parse(trainsJson)
            app.trainStore.set(trains)
        })
        app.websocket.register("addTrain", (train) => {
            app.trainStore.add(train)
        })
        app.websocket.register("removeTrain", (train) => {
            app.trainStore.remove(train)
        })

        // load all currently known trains
        app.websocket.send("getTrains")
    })
})