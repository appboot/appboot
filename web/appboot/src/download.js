export default function download(href, filename) {
  // Create a new link
  const anchor = document.createElement("a");
  anchor.href = href;
  anchor.download = filename;

  // Append to the DOM
  document.body.appendChild(anchor);

  // Trigger `click` event
  anchor.click();

  // Remove element from DOM
  document.body.removeChild(anchor);
}
