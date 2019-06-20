import { LitElement, html } from 'lit-element';

export class MatternelleElement extends LitElement {
    static get properties() {
        return {
            message: { type: String },
            pie: { type: Boolean }
        };
    }

    constructor() {
        super();
        this.initWS();
        this.loadComplete = false;
        this.message = 'Hello World from LitElement';
        this.pie = false;
    }

    initWS() {
        const socket = new WebSocket('ws://127.0.0.1:8989/ws'); //'ws://127.0.0.1:8065/plugins/com.gitlab.itk.fr.matternelle/ws'
        socket.onerror = function(error) {
            console.error(error);
        };

        // Lorsque la connexion est établie.
        socket.onopen = function(event) {
            console.log('Connexion établie.');

            // Lorsque la connexion se termine.
            this.onclose = function(event) {
                console.log('Connexion terminé.');
            };

            // Lorsque le serveur envoi un message.
            this.onmessage = function(event) {
                console.log('Message:', JSON.parse(event.data));
            };

            // Envoi d'un message vers le serveur.
            this.send(`{"command":"msg", "msg":"Hello world!"}`);
        };
    }

    render() {
        return html`
            <style>
                :host {
                    display: block;
                }
                :host([hidden]) {
                    display: none;
                }
            </style>

            <h1>Start LitElement!</h1>
            <p>${this.message}</p>

            <input
                name="myinput"
                id="myinput"
                type="checkbox"
                ?checked="${this.pie}"
                @change="${this.togglePie}"
            />

            <label for="myinput">I like pie.</label>

            ${this.pie
                ? html`
                      <lazy-element></lazy-element>
                  `
                : html``}
        `;
    }

    /**
     * Implement firstUpdated to perform one-time work on first update:
     * - Call a method to load the lazy element if necessary
     * - Focus the checkbox
     */
    firstUpdated() {
        console.log('first update');
        this.loadLazy();

        const myInput = this.shadowRoot.getElementById('myinput');
        myInput.focus();
    }

    /**
     * Event handler. Gets called whenever the checkbox fires a `change` event.
     * - Toggle whether to display <lazy-element>
     * - Call a method to load the lazy element if necessary
     */
    togglePie(e) {
        this.pie = !this.pie;
        this.loadLazy();
    }

    async loadLazy() {
        console.log('loadLazy');
        if (this.pie && !this.loadComplete) {
            return import('./lazy-element.js')
                .then(LazyElement => {
                    this.loadComplete = true;
                    console.log('LazyElement loaded');
                })
                .catch(reason => {
                    this.loadComplete = false;
                    console.log('LazyElement failed to load', reason);
                });
        }
    }
}

// Register the element with the browser
customElements.define('matternelle-element', MatternelleElement);
