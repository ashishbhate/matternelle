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

    handleClose() {
        this.state = STATE_SHOW;
    }

    handleInput(e) {
        this.msgToSend = e.target.value;
    }

    handleKeyPress(e) {
        if (e.target.value !== '') {
            if (e.key === 'Enter') {
                this.sendMsg();
            }
        }
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
                        position: fixed;
                        bottom: 20%;
                        right: 20px;
                    }
                    :host([hidden]) {
                        display: none;
                    }
                    .btn {
                        background-color: #4bda4b;
                        color: #d460db;
                        border-radius: 50%;
                        width: 100px;
                        height: 100px;
                        border: none;
                        font-weight: bold;
                        font-size: 1.2em;
                        cursor: pointer;
                    }
                </style>

                <button class="btn" @click=${this.start}>HELP</button>
            `;
        }
        return html`
            <style>
                :host {
                    display: block;
                    max-width: 500px;
                    position: fixed;
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
                .closeBtn {
                    position: absolute;
                    top: -15px;
                    right: -15px;
                    background-color: #fefefe;
                    border: 1px solid #ddd;
                    border-radius: 50%;
                    width: 30px;
                    height: 30px;
                    cursor: pointer;
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

            <button class="closeBtn" @click=${this.handleClose}>x</button>

            <ul class="listMsg">
                ${msgTemplates}
            </ul>

            <input
                type="text"
                .value=${this.msgToSend}
                @input=${this.handleInput}
                @keypress=${this.handleKeyPress}
            />
            <button @click=${this.sendMsg}>+</button>
        `;
    }
}

// Register the element with the browser
customElements.define('matternelle-element', MatternelleElement);
