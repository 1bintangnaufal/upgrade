# Day 15 - Auth & Session + Day 16 - Table Relation & File Upload

## Project Mockups

<br>

![Mockup 1](/Public/Assets/Images/Mockups/Desktop.png)
![Mockup 2](/Public/Assets/Images/Mockups/Mobile.png)
![Mockup 3](/Public/Assets/Images/Mockups/Tablet.png)

## Explanation
- Why did I combine these 2 tasks? Because task 13 is an unfinished task 14. So why not just combine them together? And I am sorry for the late push 😊
- Same reason why I combined task 15 & 16

<hr>

## Changelog

### Patch 0.15
- Fixed logout issue. everything works fine now!


### Patch 0.14
- Functioning register form,
- Functioning login form,
- Functioning image upload,
- Greetings for logged in users,
- Inaccessible 'Add', 'Edit', 'Delete' buttons if user is not logged in,
- Project data shown according to which user/author owns them,
- GIFs instead of JPGs because why not?

### Bugs
- Fixed

<br>

<hr>

## Features

### Scrollspy, etc

<br>

![Mockup 4](/Public/Assets/Images/Mockups/%232-project.png)

Instead of adding a new page for the add new project form and putting it on the nav, I think it's going to be more effective if I make it a scrollspy instead and moving the add new project form into a modal that can be triggered by the button on the right side of the project section.

![Mockup 7](/Public/Assets/Images/Mockups/%237-project-form.png)

I also added a little animation for the hamburger menu on mobile view, subtle zoom hover animations, and colored hover animation for the social media icons. Also, some icons inside several buttons for clarity.

| Desktop | Mobile |
|:---:|:---:|
|![Mockup 5](/Public/Assets/Images/Mockups/%230-main.gif)|![Mockup 6](/Public/Assets/Images/Mockups/bars.gif)|

<hr>

### Modal for Contact Form
Similar to modal for add new project form except the button trigger is on the nav just like the original task instruction.

![Mockup 7](/Public/Assets/Images/Mockups/%233-contact.png)

and of course, as the task provided, the 'Send' button will open your preferred email app with pre-filled message.

![Mockup 8](/Public/Assets/Images/Mockups/%234-email.png)

<hr>

### Tabs for Project Detail

Instead of showing the content simultaneously, I think tabs are going to be more fun! On 'Overview', we can see the image, dates, duration, icons, then when we click on 'Details', we can read through the description!

![Mockup 9](/Public/Assets/Images/Mockups/%235-prev.png)
![Mockup 10](/Public/Assets/Images/Mockups/%236-detail.png)

<hr>

### Toast

<br>

![Mockup 8](/Public/Assets/Images/Mockups/%238-toast.png)

Notice something to the bottom right corner of the screenshot? Yes. That message will appear after a successful new project submit. I want to do it too with edit and delete project but it will take time so perhaps I can improvise for that sooner or later after all the mandatory tasks.

<hr>

### Form Validations

1. Project Title

If the character input exceed the maximum number of 30 characters, an alert like this will occur:

<br>

![Mockup 9](/Public/Assets/Images/Mockups/%2310-val1.png)

<br>

2. Dates
    
    A. If the finish date is later than the any day of today at the time, it will return the provided alert.
    
    B. Also, if the start date is later than the finish date, it will return the provided alert.

<br>

![Mockup 11](/Public/Assets/Images/Mockups/%2314-val5.png)

![Mockup 10](/Public/Assets/Images/Mockups/%2313-val4.png)

<br>

3. Description

If the character input for the description field is less than the minimum of 80 characters, then an alert like this will occur:

<br>

![Mockup 12](/Public/Assets/Images/Mockups/%2311-val2.png)

<br>

4. Technologies Switch Toggles

The condition is at least one toggle must be selected (Multiple toggles can be switched on like checkboxes). If there are no toggles selected, the alert message is:

<br>

![Mockup 13](/Public/Assets/Images/Mockups/%2312-val3.png)

<br>

5. Clear Form

The button below Post button, it will empty all inputted data. It is also implemented in the contact form.

<hr>

### Edit Project Form

Aside from the delete button actually working (be careful), the edit button will redirect you to a form that will contain previously submitted data that can be immediately edited. Remind you all the form validation rules also apply here!

<br>

![Mockup 14](/Public/Assets/Images/Mockups/%2315-edit-project.png)
![Mockup 15](/Public/Assets/Images/Mockups/%2316-edit-project-val.png)

<hr>

### Responsive
(As seen on the first 3 mockups)

<hr>

## Deprecated Features (Perhaps will be back in upcoming features)

### Dark Mode
This project used to have a dark mode and it's toggle but I scrapped it temporarily because of Bootstrap complications. I'm gonna need time with that because I have to catch up with the mandatory task first. Here's the last updated project link that still has dark mode (Pure CSS) https://day-9-task.netlify.app/

### Background Pattern
Same reason and link as before.

### My Own Custom Color Palette
Same reason and link as before.

### Testimonials Page with HOF, OOP, AJAX (Javascript)
No longer necessary for the on going tasks, same link as before.

<hr>

## Upcoming Features

### Deletion Confirmation Modal and Logout Confirmation Modal

I tried to implement this feature but it messed up my database. Until I figure things out, just be careful with the delete button.

### Pagination

When certain number of project cards are reached, they will go to the next page but stay on the same route. I really want to implement this feature but I'm still figuring how.

### Auto Paragraph Indentation for Project Detail Description

When you hit enter and make a new line, after submit and you view it, it's still not seperated paragraph. I'll look for the workaround about this.

<hr>

# Thank You For Visiting and Reading!

Special thanks to https://github.com/adhxabre and my other mentors!

GIF Artists: https://giphy.com/pislices https://dribbble.com/czidler