import * as UrlPattern from 'url-pattern';

// TODO: Parse query parameters
export function matchUrl(template: string, url: string): boolean {
    const pattern = new UrlPattern(template);
    const isMatched = pattern.match(url) != null;
    return isMatched;
}
