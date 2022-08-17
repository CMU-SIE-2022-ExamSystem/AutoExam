
const downloadFile = (filename: string, text: string) => {
    let element = document.createElement('a');
    element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
    element.setAttribute('download', filename);
    element.style.display = 'none';
    document.body.appendChild(element);
    element.click();
    document.body.removeChild(element);
}

const downloadBlob = (filename: string, link_href: string) => {
    const element = document.createElement('a');
    element.href = link_href;
    element.setAttribute('download', filename);
    document.body.appendChild(element);
    element.click();
    document.body.removeChild(element);
}

export {downloadFile, downloadBlob};