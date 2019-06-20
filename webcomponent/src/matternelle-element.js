import { LitElement, html } from 'lit-element';

const STATE_HIDDEN = 'hidden';
const STATE_SHOW = 'show';
const STATE_INPUT = 'input';

export class MatternelleElement extends LitElement {
    static get properties() {
        return {
            state: { type: String }, //hidden, show, input
            msgToSend: { type: String },
            msg: { type: Array }
        };
    }

    constructor() {
        super();
        this.state = STATE_HIDDEN;
        this.msgToSend = '';
        this.msg = [];
        this.initWS();
    }

    initWS() {
        this.msgToSend = '';
        this.msg = [];
        this.socket = new WebSocket('ws://127.0.0.1:8989/ws'); //'ws://127.0.0.1:8065/plugins/com.gitlab.itk.fr.matternelle/ws'
        this.socket.onerror = error => {
            console.error(error);
        };

        this.socket.onopen = () => {
            console.log('Connexion établie.');
        };

        this.socket.onclose = () => {
            console.log('Connexion terminé.');
            setTimeout(() => {
                this.initWS();
            }, 3000);
        };

        this.socket.onmessage = event => {
            const msg = JSON.parse(event.data);
            if (
                msg.command === 'nbChatUser' &&
                msg.nbChatUser !== undefined &&
                msg.nbChatUser > 0
            ) {
                this.state = STATE_SHOW;
            }
            this.msg = [...this.msg, msg];
            console.log('Message:', msg);
        };
    }

    start() {
        this.state = STATE_INPUT;
    }

    handleInput(e) {
        this.msgToSend = e.target.value;
    }

    sendMsg() {
        const msg = { command: 'msg', msg: this.msgToSend, byAppUser: true };
        this.msg = [...this.msg, msg];
        this.socket.send(JSON.stringify(msg));
        this.msgToSend = '';
    }

    render() {
        const msgTemplates = this.msg
            .filter(i => i.command === 'msg')
            .map(
                i =>
                    html`
                        <li class="${i.byAppUser ? 'byAppUser' : ''}">
                            ${i.msg}
                        </li>
                    `
            );
        if (this.state === STATE_HIDDEN) {
            return html``;
        } else if (this.state === STATE_SHOW) {
            return html`
                <style>
                    :host {
                        display: block;
                        max-width: 500px;
                        position: absolute;
                        bottom: 20%;
                        right: 20px;
                        background-color: #fefefe;
                        border: 1px solid #ddd;
                        padding: 20px;
                        list-style: none;
                    }
                    :host([hidden]) {
                        display: none;
                    }
                </style>

                <button @click=${this.start}>START</button>
            `;
        }
        return html`
            <style>
                :host {
                    display: block;
                    max-width: 500px;
                    position: absolute;
                    bottom: 20%;
                    right: 20px;
                    background-color: #fefefe;
                    border: 1px solid #ddd;
                    padding: 20px;
                    list-style: none;
                }
                :host([hidden]) {
                    display: none;
                }
                .listMsg {
                    list-style: inherit;
                    margin: 0;
                    padding: 0;
                }
                .byAppUser {
                    text-align: right;
                }
            </style>

            <ul class="listMsg">
                ${msgTemplates}
            </ul>

            <input
                type="text"
                .value=${this.msgToSend}
                @input=${this.handleInput}
            />
            <button @click=${this.sendMsg}>+</button>
        `;
    }
}

// Register the element with the browser
customElements.define('matternelle-element', MatternelleElement);
