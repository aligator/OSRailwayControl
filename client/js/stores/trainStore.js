export default class TrainStore {
    trains = new Map()
    selected = ""

    onSpeedChange = null

    constructor(onSpeedChange) {
        this.onSpeedChange = onSpeedChange
    }

    set(trains) {
        this.trains = new Map()
        trains.forEach((train) => {
            this.trains.set(train.name, train)
        })
        this.draw()
    }

    addOrUpdate(train) {
        this.trains.set(train.name, train)
        this.draw()
        if (this.selected == train.name) {
            this.drawSelectedTrain()
        }
    }

    remove(trainName) {
        if (this.trains.has(trainName)) {
            this.trains.delete(trainName)
            if (this.selected === trainName) {
                this.selected = ""
                this.drawSelectedTrain()
            }
            this.draw()
        }
    }

    draw() {
        // draw train list
        let listHtml = "<ul>"
        this.trains.forEach((train) => {
            listHtml += `<li><button id="btn-train-list-${train.name}">${train.name}</button></li>`
        })
        listHtml += "</ul>"
        document.getElementById("train-list-container").innerHTML = listHtml

        this.trains.forEach((train) => {
            document.getElementById(`btn-train-list-${train.name}`).addEventListener("click", _ => {
                this.selected = train.name
                this.drawSelectedTrain()
            })
        })
    }

    drawSelectedTrain() {
        if (!this.selected) {
            document.getElementById(`selected-train`).innerHTML = ""
            return
        }

        const train = this.trains.get(this.selected)
        let html = `
<h2>${this.selected}</h2>
<button id="btn-backward">backward</button><button id="btn-stop">stop</button><button id="btn-forward">forward</button>
<div class="slide-container">
  <label for="slider-speed">Speed (0-1023)</label>
  <input type="range" min="0" max="1023" value=${train.speed} class="slider" id="slider-speed">
  <div>${train.speed}</div>
</div>`
        if (!this.onSpeedChange) {
            return
        }

        document.getElementById(`selected-train`).innerHTML = html

        document.getElementById(`btn-forward`).addEventListener("click", _ => {
            train.direction = 1
            this.onSpeedChange(this.selected, train.direction * train.speed)
        })

        document.getElementById(`btn-backward`).addEventListener("click", _ => {
            train.direction = -1
            this.onSpeedChange(this.selected, train.direction * train.speed)
        })

        document.getElementById(`btn-stop`).addEventListener("click", _ => {
            train.direction = 0
            this.onSpeedChange(this.selected, 0)
        })

        document.getElementById(`slider-speed`).addEventListener("change", e => {
            train.speed = e.target.valueAsNumber
            this.drawSelectedTrain()
            this.onSpeedChange(this.selected, train.direction * train.speed)
        })
    }
}