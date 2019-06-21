import { customElement, html, LitElement, property } from 'lit-element';

const STATE_HIDDEN = 'hidden';
const STATE_SHOW = 'show';
const STATE_INPUT = 'input';

@customElement('matternelle-element')
export class MatternelleElement extends LitElement {
  private socket: WebSocket | null = null;
  private state: 'hidden' | 'show' | 'input' = STATE_HIDDEN;
  private msgToSend: string = '';
  private msg: Array<any> = [];

  @property({ type: String, reflect: true })
  user: string = '';

  @property({ type: String, reflect: true })
  tokenApp: string = '';

  constructor() {
    super();
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
      if (this.tokenApp) {
        this.socket.send(
          JSON.stringify({ command: 'tokenApp', tokenApp: this.tokenApp })
        );
      }
    };

    this.socket.onclose = () => {
      console.log('Connexion terminé.');
      setTimeout(() => {
        this.initWS();
      }, 3000);
    };

    this.socket.onmessage = event => {
      const msg = JSON.parse(event.data);
      console.log('Message:', msg);
      if (
        msg.command === 'nbChatUser' &&
        msg.nbChatUser !== undefined &&
        msg.nbChatUser > 0
      ) {
        this.state = STATE_SHOW;
        if (this.user) {
          this.socket.send(JSON.stringify({ command: 'msg', msg: this.user }));
        }
      }
      this.msg = [...this.msg, msg];
      this.requestUpdate();
    };
  }

  attributeChangedCallback(name, oldval, newval) {
    console.log('attribute change: ', name, newval);
    if (name === 'user' && this.state !== STATE_HIDDEN) {
      this.socket.send(JSON.stringify({ command: 'msg', msg: this.user }));
    } else if (name === 'token' && this.state !== STATE_HIDDEN) {
      JSON.stringify({ command: 'tokenApp', tokenApp: this.tokenApp });
    }
    super.attributeChangedCallback(name, oldval, newval);
  }

  start() {
    this.state = STATE_INPUT;
    this.requestUpdate();
  }

  handleClose() {
    this.state = STATE_SHOW;
    this.requestUpdate();
  }

  handleInput(e) {
    this.msgToSend = e.target.value;
    this.requestUpdate();
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
    this.requestUpdate();
  }

  protected render() {
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
