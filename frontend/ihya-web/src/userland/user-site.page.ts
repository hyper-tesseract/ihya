import { css, html } from 'lit';
import { customElement } from 'lit/decorators.js';
import { Block } from '../ihya-lib/block.component';
import './user-site-root.component';

type RouteInfo = {
    subdomain: string;
    path: string;
};

@customElement('ihya-site-wrapper')
export class IhyaSiteWrapper extends Block {
    declare author: string;

    declare portfolio_url: string;

    routeInfo: RouteInfo;

    constructor() {
        super();
        this.routeInfo = IhyaSiteWrapper.getSubdomainFromUrl(
            window.location.href
        );
        this.author = 'Muhammad Azeem @TangentLabs';
        this.portfolio_url = 'https://github.com/tangent-cloud/';
    }

    ihyaFooter() {
        return html`
            <div class="wrapper-footer">
                Made by
                <a
                    target="_blank"
                    rel="noopener noreferrer"
                    href=${this.portfolio_url}
                    >${this.author}</a
                >
            </div>
        `;
    }

    render() {
        return html`
            <user-site-root
                persona_id=${this.routeInfo.subdomain}
                page=${this.routeInfo.path}
            ></user-site-root>
            ${this.ihyaFooter()}
        `;
    }

    static getSubdomainFromUrl(url: any): RouteInfo {
        const parsedUrl = new URL(url);
        const subdomain = parsedUrl.hostname.split('.')[0]; // Assuming only one subdomain
        const path = parsedUrl.pathname;
        return { subdomain, path };
    }

    static styles = css`
        .wrapper-footer {
            position: relative;
            bottom: 0;
            left: 0;
            right: 0;
            top: 0;
            background-color: #222; /* Dark metallic grey */
            color: #fff;
            font-size: calc(12px + 0.5vmin);
            display: flex;
            padding: 10px;
            margin: 0px;
        }

        .wrapper-footer a {
            margin-left: 5px;
            color: inherit; /* Inherit the font color from the .app-footer element */
            text-decoration: none;
        }

        .wrapper-footer a:hover {
            text-decoration: underline; /* Underline on hover */
            color: #3399ff; /* Accent color for hover (bright blue) */
        }
    `;
}
