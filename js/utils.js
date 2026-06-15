// Navigate to a page
function goTo(page) {
  window.location.href = page;
}

// Show toast then redirect
function showToastAndRedirect(toastId, redirectTo, delay = 1500) {
  const toast = document.getElementById(toastId);
  if (toast) {
    toast.style.display = 'block';
    setTimeout(() => { window.location.href = redirectTo; }, delay);
  }
}

// Keyboard helpers
function typeKey(inputId, key) {
  const el = document.getElementById(inputId);
  if (!el) return;
  el.value += key;
  el.focus();
}

function deleteKey(inputId) {
  const el = document.getElementById(inputId);
  if (!el) return;
  el.value = el.value.slice(0, -1);
  el.focus();
}
