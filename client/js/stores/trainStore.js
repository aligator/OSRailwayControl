export default class TrainStore {
    trains = new Set()

    selected = ""

    onSpeedChange = null

    constructor(onSpeedChange) {
        this.onSpeedChange = onSpeedChange
    }

    set(trains) {
        this.trains = new Set()
        trains.forEach((train) => {
            this.trains.add(train)
        })
        this.draw()
    }

    add(train) {
        this.trains.add(train)
        this.draw()
    }

    remove(train) {
        if (this.trains.has(train)) {
            this.trains.delete(train)
            if (this.selected === train) {
                this.selected = ""
            }
            this.draw()
        }
    }

    draw() {
        // draw train list
        let listHtml = "<ul>"
        this.trains.forEach((train) => {
            listHtml += `<li><button id="btn-train-list-${train}">${train}</button></li>`
        })
        listHtml += "</ul>"
        document.getElementById("train-list-container").innerHTML = listHtml

        this.trains.forEach((train) => {
            document.getElementById(`btn-train-list-${train}`).addEventListener("click", _ => {
                this.selected = train
                this.drawSelectedTrain()
            })
        })
    }

    drawSelectedTrain() {
        if (!this.selected) {
            document.getElementById(`selected-train`).innerHTML = ""
            return
        }

        let html = `<h2>${this.selected}</h2><button id="btn-backward">backward</button><button id="btn-stop">stop</button><button id="btn-forward">forward</button>`
        if (!this.onSpeedChange) {
            return
        }

        document.getElementById(`selected-train`).innerHTML = html

        document.getElementById(`btn-forward`).addEventListener("click", _ => {
            this.onSpeedChange(this.selected, 1023)
        })

        document.getElementById(`btn-backward`).addEventListener("click", _ => {
            this.onSpeedChange(this.selected, -1023)
        })

        document.getElementById(`btn-stop`).addEventListener("click", _ => {
            this.onSpeedChange(this.selected, 0)
        })
    }
}