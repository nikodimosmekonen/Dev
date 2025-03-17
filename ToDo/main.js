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
    const tasks = new Map(); // Initialize tasks array
  
    button.addEventListener("click", function() {
      const taskText = input.value.trim(); // Trim whitespace
      if (taskText !== "") { // Prevent adding empty tasks
          const id= new Date().toISOString();
          tasks.set(id ,taskText);
          const newParagraph = document.createElement("p"); // Create <p> element
          newParagraph.textContent = tasks.get(id); // Set its content (more efficient)
          newParagraph.setAttribute("class","tasks")
          newParagraph.setAttribute("onclick","edit_task(this)")
          newParagraph.setAttribute("id",id)
          taskDiv.appendChild(newParagraph); // Append <p> to taskDiv
          input.value = ""; // Clear the input field
      }
    });
    
  });
  function edit_task(element){
    const input = prompt("EDIT TASK",element.textContent)
    element.innerHTML=input
    
  }
