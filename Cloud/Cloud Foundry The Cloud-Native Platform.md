Cloud Foundry The Cloud-Native Platform
=======================================


## 1. Introduction
- To thrive in this world, you need the ability to deliver software quickly, repeatedly, and with regular feedback.
- Cloud Foundry is a platform for running applications and services.
- Its purpose is to change the way applications and services are deployed and run by reducing the develop to deployment cycle time.

### The Competitive Advantage
- Market-disrupting companies differ from incumbents because they can repeatedly deliver software,  
    with velocity, through iterative development cycles of short duration.

### The Development-Feedback Cycle
- Any development cycle must include a constant feedback loop from end users for continually refining your product.
- A constant developmentfeedback cycle allows you to  
    try out new ideas, quickly identify failings, acknowledge feedback, rapidly adapt, and repeat.

### Velocity Over Speed
- In every marketplace, speed wins.
- Speed coupled with the development-feedback cycle produces velocity, and velocity is the only way to move both quickly and securely.

### The Critical Challenge
- Cloud-native platforms provide a focal point that centers that change.  
    They make it easier to do the right thing because everything you need to deliver software is already built into a platform.

### Becoming Cloud Native
- Cloud native is a term describing software designed to run and scale reliably and predictably  
    on top of potentially unreliable cloud-based infrastructure.
- Cloud-native applications are purposefully designed to be infrastructure unaware,  
    meaning they are decoupled from infrastructure and free to move as required.
- Becoming cloud native involves three fundamental tenets:
    - Automated Infrastructure Management and Orchestration
    - Platforms
    - The Twelve-Factor App
- Becoming cloud native is an essential step in establishing a timely development-feedback cycle  
    because it helps companies achieve velocity deploying software releases into production.

### Undifferentiated Heavy Lifting
- Most enterprises outside the information technology industry do not generate revenue from selling software;  
    they leverage software to drive value into their core business.
- Examples of undifferentiated heavy lifting include:
    - Provisioning VMs, middleware, and databases
    - Creating and orchestrating containers
    - User management
    - Load balancing and traffic routing
    - Centralized log aggregation
    - Scaling
    - Security auditing
    - Providing fault tolerance and resilience

### Platforms Benefit Developers
- The platform leverages middleware and infrastructure directly, allowing streamlined development through self-service environments.
- Applications can then be bound to a wide set of available backing services that are available on demand.

### Platforms Benefit Operations
- The platform provides responsive IT operations,  
    with full visibility and control over the application life cycle, provisioning, deployment, upgrades, and security patches.

### Platforms Benefit the Business
- The business no longer needs to be constrained by process or organizational silos.
- Cloud-native platforms provide a contractual promise to allow the business to move with velocity and establish the developer-feedback loop.

### Chapter Summary
- Technology is used to achieve a competitive advantage, but technology alone has never been enough.
- The world needs to fundamentally change the way it builds and deploys software  
    in order to succeed in the hugely competitive markets that exist today.
- Cloudnative platforms provide the most compelling way to enable that fundamental shift.


## 2. Adapt or Die
> There are two approaches to handling change: adapt or die vs. same mess for less!
- Businesses today are constantly pressured to adopt the myriad of technical driving forces impacting software development and delivery.  
    These driving forces include:
    - Anything as a service
    - Cloud computing
    - Containers
    - Agile
    - Automation
    - DevOps
    - Microservices
    - Business-capability teams
    - Cloud-native applications

### Anything As A Service
- We have now reached the point where if you are not leveraging compute resources as a service,  
    it is unlikely that you are moving at the pace required to stay competitive.
- If there is a service out there that has been deployed and managed in a repeatable and scalable way,  
    becoming a consumer of that service allows you to focus on software supporting your revenue-generating business.

### Cloud Computing
- Cloud computing is the third incarnation of the platform eras:
    - The first era was the mainframe, which dominated the industry from the 1950s through the early 1980s.
    - Client-server architecture was established as the dominant second platform from the 1980s right up until the second decade of the 21st century.
    - The “as a service” movement in IT has broadly been termed cloud computing, and it defines the third platform era we live in today.
- Cloud computing has been subdivided into “as a service” layers (SaaS, PaaS, IaaS, etc).
- Regardless of the layer, there are three definitive attributes of “as a service”:
    - Elasticity  
        The ability to handle concurrent growth through dynamically scaling the service up and down at speed.
    - On demand  
        The ability to choose when and how to consume the required service.
    - Self-service  
        The ability to directly provision or obtain the required service without time-consuming ticketing.
- This self-service capability is a shift from procuring resources through a ticketing system involving handoffs  
    and delays between developers and operations.

### Platform as a Service
- IaaS and SaaS are generally well understood concepts.
- SaaS provides the ability to consume software services on demand without having to procure, license, and install binary packages.
- A cloud-native platform describes a platform designed to reliably and predictably run and scale on top of potentially unreliable cloudbased infrastructure.

### Containers
- In recent years, there has been a rapid growth in container-based technologies (such as LXC, Docker, Garden, and Rocket).
- Containers offer three distinct advantages over traditional VMs:
    1. Speed and efficiency
    2. Greater resource consolidation
    3. Application stack portability
- they have enabled a new era of application stack portability because  
    applications and dependencies developed in a container can be easily moved and run in different environments.

### Understanding Containers
- Containers are best understood as having two elements:
    - Container images: These package a repeatable runtime environment (encapsulating your application, dependencies, and file system)  
        in a way that allows for images to be moved between hosts.  
        Container images have instructions that specify how they should be run but they are not explicitly self-executable,  
        meaning they cannot run without a container management solution.
    - A container management solution: This often uses kernel features such as Linux namespaces to run a container image in isolation,  
        often within a shared kernel space.  
        Container management solutions arguably have two parts: 
            - The frontend management component known as a container engine such as Dockerengine or Garden-Linux
            - The backend container runtime such as runC or runV.

### Agile
- Agile software development can best be understood by referring to the Agile Software Development Manifesto.
- This manifesto values:
    1. Individuals and interactions over processes and tools
    2. Working software over comprehensive documentation
    3. Customer collaboration over contract negotiation
    4. Responding to change over following a plan
- The Agile methodology is an alternative to traditional sequential development strategies, such as the waterfall approach.
- Most teams now define epics that are broken down into smaller user stories weighted by a point system.
- Stories are implemented over sprints with inception planning meetings,  
    daily standups, and retrospective meetings to showcase demonstrable functions to key stakeholders.
- Agile deployment allows teams to test out new ideas, quickly identify failings with rapid feedback, learn, and repeat.

### Automate Everything
- Operational practices around continuous integration (CI) and continuous delivery (CD) have been established  
    to address the following two significant pain points with software deployment:
    - Handoffs between different teams cause delays.  
        Handoffs can occur during several points throughout an application life cycle, starting with procuring hardware,  
        right through to scaling or updating applications running in production.
    - Release friction describes the constant need for human interaction as opposed to using automation for releasing software.
- When establishing deployment pipelines, it is important to understand the progression of continuous integration and continuous delivery.

#### Continuous Integration
- Continuous integration (CI) is a development practice.  
    Developers check code into a central shared repository.  
    Each check-in is verified by automated build and testing, allowing for early detection problems and consistent software releases.
- Continuous integration has enabled streamlined efficiencies for the story to demo part of the cycle.
- continuous integration without continuous delivery into production means  
    you only have a small measure of agility confined by a traditional waterfall process.

#### Continuous Delivery
- Continuous delivery further extends continuous integration.  
    The output of every code commit is a release candidate that progresses through an automated deployment pipeline  
    to a staging environment, unless it is proven defective.
- If tests pass, the release candidate is deployable in an automated fashion.
- Companies that operate in this way have a significant advantage and are able to create products that constantly adapt to feedback and user demands.

### DevOps
- DevOps is a software development and operations culture that has grown in popularity over the recent years.
- The method acknowledges the interdependence of:
    - Software development
    - Quality assurance and performance tuning
    - IT operations
    - Administration (SysAdmin, DBAs, etc.)
    - Project and release management
- DevOps aims to span the full application life cycle to help organizations rapidly produce  
    and operationally maintain performant software products and services.
- Lack of overall ownership and constant handoffs from one department to the next can prolong a task from minutes or hours to days or weeks.
- Shared toolsets and a common language are established around the single goal of developing and supporting software running in production.
- Silos are replaced by collaboration with all members under the same leadership.
- They develop products instead of working on projects. Products, if successful, are long lived.

### Microservices
- Microservices is a term used to describe a software architectural style that has emerged over the last few years.
- It describes a modular approach to building software in which complex applications are composed of several small,  
    independent processes communicating with each other though explicitly defined boundaries using language-agnostic APIs.
- These smaller services focus on doing a single task very well. They are highly decoupled and can scale independently.

P. 17
