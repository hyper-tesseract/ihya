import './userland/user-content/themes/solid-state-html5up/component/navbar.component';
import './userland/user-content/themes/solid-state-html5up/element/footer.segment';
import './userland/user-site.page';

import { html } from 'lit';
import { IhyaPage } from './ihya-lib/base.page';

class App extends IhyaPage {
    public render() {
        return html`<ihya-site-wrapper></ihya-site-wrapper>`;
    }
}

customElements.define('ihya-app', App);

function loadThemeStyleSheet(cssUrl: string): void {
    const linkElement = document.createElement('link');
    linkElement.rel = 'stylesheet';
    linkElement.href = cssUrl;
    document.head.appendChild(linkElement);
}

loadThemeStyleSheet(
    'http://localhost:5000/theme/solid-state-html5up/css/main.css'
);
