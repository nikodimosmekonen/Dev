const tasks = new Map(); // Global task storage

window.addEventListener('load', () => {
  const taskDiv = document.getElementById('taskDiv');
  const addButton = document.getElementById('add');
  const taskInput = document.getElementById('task_input');

  // Event delegation for edit clicks
  taskDiv.addEventListener('click', (e) => {
    if (e.target.classList.contains('tasks_label')) {
      enterEditMode(e.target);
    }
  });

  addButton.addEventListener('click', () => {
    const text = taskInput.value.trim();
    if (text) {
      const id = Date.now().toString();
      tasks.set(id, text);
      renderTask(id, text);
      taskInput.value = '';
    }
  });
});

function renderTask(id, text) {
  const li = document.createElement('li');
  li.className = 'tasks_li';
  
  const checkbox = document.createElement('input');
  checkbox.type = 'checkbox';

  const label = document.createElement('label');
  label.className = 'tasks_label';
  label.dataset.taskId = id; // Use data attribute instead of ID
  label.textContent = text;

  li.append(checkbox, label);
  document.getElementById('taskDiv').appendChild(li);
}

function enterEditMode(label) {
  const id = label.dataset.taskId;
  const originalText = tasks.get(id);

  const form = document.createElement('form');
  form.innerHTML = `
    <input type="text" 
           id="edit_${id}" 
           name="task_edit" 
           value="${originalText}"
           class="edit-input">
    <button type="submit">Update</button>
  `;

  // Replace label content with form
  label.innerHTML = '';
  label.appendChild(form);

  // Focus input immediately
  form.querySelector('input').focus();

  // Handle form submission
  form.addEventListener('submit', (e) => {
    e.preventDefault();
    const newText = form.querySelector('input').value.trim();
    if (newText) {
      tasks.set(id, newText);
      label.textContent = newText; // Revert to label state
    }
  });
}