/* input = document.getElementById("task_input")
button = document.getElementById("add")
taskDiv= document.getElementById("taskDiv")
tasks=[]
     
button.addEventListener("click", function() {
    tasks.push(input.value);
    let para = document.createElement("div")
    para.id = "tasks"
    let text = document.createTextNode(input.value)
    para.appendChild(text)
    taskDiv.appendChild(para)
    input.value = "" // clear input field after adding a task
}); */
// Wait for the DOM to be fully loaded before accessing elements

window.addEventListener('load', function() {
    const input = document.getElementById("task_input");
    const button = document.getElementById("add");
    const taskDiv = document.getElementById("taskDiv");
    const tasks = []; // Initialize tasks array
  
    button.addEventListener("click", function() {
      const taskText = input.value.trim(); // Trim whitespace
      if (taskText !== "") { // Prevent adding empty tasks
          tasks.push(taskText);
          const newParagraph = document.createElement("p"); // Create <p> element
          newParagraph.textContent = taskText; // Set its content (more efficient)
          taskDiv.appendChild(newParagraph); // Append <p> to taskDiv
          input.value = ""; // Clear the input field
      }
    });
  });
