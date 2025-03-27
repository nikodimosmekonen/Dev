const nav_list = document.querySelector('.nav-list');
const mobile_nav_toggle = document.querySelector('.mobile-nav-toggle');

mobile_nav_toggle.addEventListener('click',(e)=>{
    if(nav_list.getAttribute("data-visible") === "false"){
        nav_list.setAttribute("data-visible",true);
        mobile_nav_toggle.setAttribute("aria-expanded",true);
    }else{
        nav_list.setAttribute("data-visible",false);
        mobile_nav_toggle.setAttribute("aria-expanded",false)
    }
})
