import { LitElement, html } from 'lit-element';

export class MatternelleElement extends LitElement {
    static get properties() {
        return {
            enable: { type: Boolean },
            msgToSend: { type: String },
            msg: { type: Array }
        };
    }

    constructor() {
        super();
        this.enable = false;
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
                this.enable = true;
            }
            this.msg = [...this.msg, msg];
            console.log('Message:', msg);
        };
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
        if (!this.enable) {
            return html``;
        }
        return html`
            <style>
                :host {
                    display: block;
                }
                :host([hidden]) {
                    display: none;
                }
                .byAppUser {
                    text-align: right;
                }
            </style>

            <ul>
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
