export default class TrainStore {
    trains = new Map()

    /**
     * selected contains the name of the selected train
     * @type {string}
     */
    selected = ""

    /**
     * patchTrain is a callback which changes some fields of the train.
     * You have to pass only the fields you want to change.
     * The callback gets the train name and the trainFields as parameters.
     * trainFields has to be an object which has at least one of the fields set.
     *
     * @type {(trainName: string; trainFields: {
     *     speed?: {string},
     *     headlights?: {boolean},
     *     backlights?: {boolean}
     * }) => void}
     */
    patchTrain = null

    constructor(onSpeedChange) {
        this.patchTrain = onSpeedChange
    }

    /**
     * Set all trains, replaces the existing ones.
     * @param trains
     */
    set(trains) {
        this.trains = new Map()
        trains.forEach((train) => {
            this.trains.set(train.name, train)
        })
        this.drawTrainList()
    }

    /**
     * addOrUpdate one train. If does not exist yet, it is just added else it gets updated.
     */
    addOrUpdate(train) {
        this.trains.set(train.name, train)
        this.drawTrainList()
        if (this.selected == train.name) {
            this.drawTrainControl()
        }
    }

    /**
     * remove deletes a train by the given name.
     * @param trainName
     */
    remove(trainName) {
        if (this.trains.has(trainName)) {
            this.trains.delete(trainName)
            if (this.selected === trainName) {
                this.selected = ""
                this.drawTrainControl()
            }
            this.drawTrainList()
        }
    }

    drawTrainList() {
        let listHtml = ""

        // order list by name
        Array.from(this.trains.values()).sort((a, b) => {
            if (a.name === b.name) {
                return 0;
            } else if (a.name > b.name) {
                return 1
            } else {
                return -1
            }
        }).forEach((train) => {
            // add a button for each train
            listHtml += `<button type="button" class="btn btn-primary btn-train-list ${train.name === this.selected ? "active" : ""}" id="btn-train-list-${train.name}">${train.name}</button>`
        })

        document.getElementById("train-list-container").innerHTML = listHtml

        // assign a listener to each button
        this.trains.forEach((train) => {
            document.getElementById(`btn-train-list-${train.name}`).addEventListener("click", _ => {
                this.selected = train.name
                // redraw train list to activate the correct button
                this.drawTrainList()
                // draw train control for the selected train
                this.drawTrainControl()
            })
        })
    }

    /**
     * drawTrainControl redraws the whole train control page
     */
    drawTrainControl() {
        // clear train control if no train is selected
        if (!this.selected) {
            document.getElementById(`selected-train`).innerHTML = ""
            return
        }

        const train = this.trains.get(this.selected)
        let html = `
<h2>${this.selected}</h2>
<button type="button" class="btn btn-primary btn-train-control" id="btn-backward">backward</button><button type="button" class="btn btn-danger btn-train-control" id="btn-stop">stop</button><button type="button" class="btn btn-primary btn-train-control" id="btn-forward">forward</button>
<div class="slide-container">
  <label for="slider-speed">Change speed (0-1023):</label><br>
  <input type="range" min="0" max="1023" value=${train.speed} class="slider" id="slider-speed">
  <div>Resulting speed: ${train.direction * train.speed}</div><br>
  <label for="chk-headlights">Enable headlights:</label>
  <input type="checkbox" ${train.headlights ? "checked" : ""} id="chk-headlights"><br>
  <label for="chk-backlights">Enable backlights:</label>
  <input type="checkbox" ${train.backlights ? "checked" : ""} id="chk-backlights"><br>
</div>`
        if (!this.patchTrain) {
            return
        }
        document.getElementById(`selected-train`).innerHTML = html

        document.getElementById(`btn-forward`).addEventListener("click", _ => {
            train.direction = 1
            this.drawTrainControl()
            this.patchTrain(this.selected, {
                speed: train.speed * train.direction
            })
        })

        document.getElementById(`btn-backward`).addEventListener("click", _ => {
            train.direction = -1
            this.drawTrainControl()
            this.patchTrain(this.selected, {
                speed: train.speed * train.direction
            })
        })

        document.getElementById(`btn-stop`).addEventListener("click", _ => {
            train.speed = 0
            this.drawTrainControl()
            this.patchTrain(this.selected, {
                speed: train.speed * train.direction
            })
        })

        document.getElementById(`slider-speed`).addEventListener("change", e => {
            train.speed = e.target.valueAsNumber
            this.drawTrainControl()
            this.patchTrain(this.selected, {
                speed: train.speed * train.direction
            })
        })

        document.getElementById(`chk-headlights`).addEventListener("change", e => {
            train.headlights = !train.headlights
            this.drawTrainControl()
            this.patchTrain(this.selected, {
                headlights: train.headlights
            })
        })

        document.getElementById(`chk-backlights`).addEventListener("change", e => {
            train.backlights = !train.backlights
            this.drawTrainControl()
            this.patchTrain(this.selected, {
                backlights: train.backlights
            })
        })
    }
}