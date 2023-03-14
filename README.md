- [Preview](#preview)
- [Highlight](#highlight)
    - [Facebook Smart Traffic And LB](#facebook-smart-traffic)
    - [How Facebook Disappeared From The Internet And Down](#how-facebook-disappeared-from-the-internet-and-down)
    - [All Of Network And Ten Minute](#all-of-network-in-ten-minute)

## Preview <a name="preview"></a>

When we develop one function for mini project, it hasn't complex requirements for the system. Usually, we define the
feature, define logic and implement. </br>

It is also that feature, when it is deployed on a big product, a lot of problems arise beyond making sure it runs
correct. Ex: load capacity, high performance, resource optimization, DB and store is good when run long time, security,
money for this function, ... </br>

This problem is not simple, you need multi knowledge and skills to handle that, including three important skills:
infra + backend + clear requirement + security. You need a top-down, non-function to function, design system to detail,
after is implementation. </br>  </br>

In this document, I will research and survey the infra, devops, backend components of the project with user sizes
ranging from a few hundred thousand to about five hundred million users and more. All problems and solutions will be
described in detail from theory, implementation, code, testing on product environment. I hope this is an advanced and
practical document for me and someone to improve their skills in the web app field. I'm very happy if that comes true,
happy and relax with it. </br>

## Highlight <a name="highlight"></a>
## Facebook Smart Traffic And LB <a name="facebook-smart-traffic"></a>
