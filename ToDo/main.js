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
  const div = document.createElement('div');
  div.className = 'tasks_div';
  
  const checkbox = document.createElement('input');
  checkbox.type = 'checkbox';

  const label = document.createElement('label');
  label.className = 'tasks_label';
  label.dataset.taskId = id; 
  label.textContent = text;

  div.append(checkbox, label);
  document.getElementById('taskDiv').appendChild(div);
}

function enterEditMode(label) {
  const id = label.dataset.taskId;
  const originalText = tasks.get(id);

  const edit_div = document.createElement('div');
  edit_div.innerHTML = `
    <input type="text" 
           id="edit_${id}" 
           name="task_edit" 
           value="${originalText}"
           class="edit-input">
    <button onclick="Update('${id}')">Update</button>
    <button onclick="Delete('${id}')">Delete</button>
  `;

  // Replace label content with form
  label.innerHTML = '';
  label.appendChild(edit_div);

  // Focus input immediately
  edit_div.querySelector('input').focus();
}
function Update(id){
  const label = document.querySelector(`[data-task-id="${id}"]`);
  const input = document.querySelector(`[id="edit_${id}"]`);
  const newText = input.value.trim();
  if (newText) {
    tasks.set(id, newText);
    label.innerHTML = '';
    label.textContent = newText;
    }
}
function Delete (id){
  const div = document.querySelector(`[data-task-id="${id}"]`).closest("div");
  div.innerHTML = '';
  tasks.delete(id)
  div.remove()
}