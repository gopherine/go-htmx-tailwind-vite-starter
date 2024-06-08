import * as _ from 'lodash';

function initialize() {
  console.log("loadash.ts: initialize function called");
  const output = document.getElementById('output');
  console.log("loadash.ts: Output element:", output);
  if (output) {
    output.textContent = _.join(['Hello', 'World'], ' ');
  } else {
    console.error("loadash.ts: Output element not found");
  }
}

// Listen for the DOMContentLoaded event
window.addEventListener('DOMContentLoaded', initialize);

// Export the initialize function as default
export default initialize;