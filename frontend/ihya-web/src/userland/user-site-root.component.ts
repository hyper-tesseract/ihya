import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';
import './user-content/themes/solid-state-html5up/portfolio.page';

type UserSiteContent = {
    header_content: any;
    page_content: any;
    footer_content: any;
};

@customElement('user-site-root')
export class UserSiteRoot extends LitElement {
    private data: UserSiteContent | undefined;

    static get properties() {
        return {
            persona_url: { type: String },
            page: { type: String },
        };
    }

    fetchData(): Promise<any> {
        return fetch('https://localhost:5000/api/user-site/')
            .then(response => response.json())
            .then(data => {
                this.data = data;
            })
            .catch(error => {
                console.error('Error fetching data: %s', error);
            });
    }

    // connectedCallback() {
    //   super.connectedCallback();
    //   this.fetchData().then(() => this.requestUpdate());
    // }

    render() {
        return html`<solid-state-portfolio-page></solid-state-portfolio-page>`;
    }
}
