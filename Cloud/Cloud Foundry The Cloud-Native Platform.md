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
- Teams are organized around structuring software in smaller chunks (features, not releases)  
    as separate deployments that can move and scale independently.

### Business-Capability Teams
- if you want your architecture to be centered around business capability instead of specialty,  
    you should go all in and structure your organization appropriately to match the desired architecture.
- Once you have adopted this approach, the next step is to define what business-capability teams are needed.

### Cloud-Native Applications
- An architectural style known as cloud-native applications has been established to describe the design of applications  
    specifically written to run in a cloud environment.
- The [Twelve Factor App](https://12factor.net/) explains the 12 principles underpinning cloudnative applications.
- Twelve Factor can be thought of as the contract between an application and a cloud-native platform.
- Similarly, platform contracts are born out of previous tried and tested constraints;  
    they are enabling and they make doing the right thing easy for developers.

### Chapter Summary
- There has been a systemic move to consuming services beyond simply provisioning VMs on demand.  
    Consuming services allows you to focus on building business software against the highest level of abstraction possible.
- For any company to be disruptive through software, it starts with the broad and complete adoptions of IaaS for compute resource to provide on-demand,  
    elastic, and self-service benefits.
- Platforms describe the layer sitting above the IaaS layer, leveraging it directly by  
    providing an abstraction of infrastructure services and a contract for applications and backing services to run and scale.
- Recently there has been a rapid growth in container-based solutions because they offer isolation, resource reuse, and portability.
- Agile development coupled with continuous delivery provides the ability to deploy new functionality at will.
- The DevOps movement has broken down the organizational silos and empowered teams to do what is right during  
    application development, deployment, and ongoing operations.
- Software should be centered around business-capability teams instead of specialty capabilities.  
    This allows for a more modular approach to building microservices software with decoupled and well defined boundaries.  
    Microservices that have been explicitly designed to thrive in a cloud environment have been termed cloud-native applications.


## 3. Cloud-Native Platforms
> Cloud Foundry is so resilient that the reliability of the underlying infrastructure becomes inconsequential.

### You Need a Cloud-Native Platform, Not a PaaS
- Cloud Foundry is a cloud-native platform offering features that can be consumed “as a service.” Historically,  
    it has been referred to as a platform as a service (PaaS).

#### Legacy PaaS
- Compared to the IaaS or SaaS layers, PaaS is not well understood because it is an overloaded and ambiguous acronym causing confusion.
- Early versions of PaaS struggled to gain broad market adoption because of:
    - Limitations around visibility into the platform
    - Lack of integration points to extend the platform
    - Limited or single language / framework support
    - Lock-in concern (due to no open source ecosystem) with a lack of non-public cloud options
- The term PaaS should die, if it is not dead already. However,  
    the reality is that the term PaaS is still out there in the marketplace, and its usage may not become obsolete as fast as it should.

#### Cloud-Native Platforms
- Cloud Foundry is an opinionated, structured platform that rectifies the PaaS confusion by imposing a strict contract between:
    - The infrastructure layer underpinning it
    - The applications and services that it supports
- Cloud-native platforms offer a super set of functionality over and above the earlier PaaS offerings.
- Their inbuilt features, such as resiliency, log aggregation, user management and security
- Cloud-native platforms are focused on what they enable you to achieve,  
    meaning what is important is not so much what a platform is, or what it does, but rather what it enables you to achieve.
- It has the potential to make the software build, test, deploy, and scale cycle significantly faster.
- The Cloud Foundry cloud-native platform has three defining characteristics: it is structured, opinionated, and open.

##### The Structured Platform
- Within the platform space, two distinct architectural patterns have emerged: structured and unstructured:
    - Structured platforms provide built-in capabilities and integration points for key concerns such as enterprise-wide user management,  
        security, and compliance.  
        Everything you need to run your applications should be provided in repeatable way, regardless of what infrastructure you run on.  
        Cloud Foundry is a perfect example of a structured platform.
    - Unstructured platforms have the flexibility to define a bespoke solution at a granular level.  
        An example of an unstructured platform would involve a “build your own platform” approach with a mix of cloud-provided services  
        and homegrown tools, assembled for an individual company.
- Structured platforms are focused on eliminating the earlier PaaSrelated problems mentioned above. For example, Cloud Foundry provides:
    - A rich and continuous stream of log information  
        with integration points into application performance management (APM) solutions for visibility into how applications are performing
    - A rich set of continually streamed metrics for understanding how the platform itself is operating
    - Integration points into existing enterprise technologies (database services, message brokers, LDAP, SAML, etc.)
    - Support for numerous languages, frameworks, and services (polyglot)
    - An open source code base with a large supporting ecosystem, backed by a foundation of over 60 companies and an API not tied to any one vendor
    - Support for a number of different deployment options, including public and non-public cloud infrastructure
- Unstructured platforms can appeal to startups, as they are typically suited to pure greenfield development with no legacy IT applications or technical debt.

##### Platform Opinions
- Opinions produce contracts to ensure applications are constrained to do the right thing.
- Like frameworks, which became popular in the early 2000s,  
    platforms are opinionated because they make specific assumptions and optimizations to remove complexity and pain from the user.
- Platforms should have opinions on how your software is deployed, run, and scaled, not where an application is deployed;  
    this means that, with respect to infrastructure choice, applications should run anywhere.

##### The Open Platform
- Cloud Foundry is an open platform. It is open on three axes:
    1. It allows a choice of IaaS layer to underpin it (AWS, vSphere, etc.).
    2. It allows for a number of different developer frameworks, polyglot languages, and application services (Ruby, Go, Spring, etc.).
    3. It is open sourced under an Apache 2 license and governed by a multi-organization foundation.

### Choice of Infrastructure
- Cloud Foundry is designed to leverage your infrastructure of choice. As such, it has been referred to as the operating system of the cloud.
- Cloud Foundry abstracts the infrastructure’s compute resource (specifically virtual storage, networking, RAM, and CPU).
- Applications can be freely moved between environments without complicated refactoring of the application or service layer.

### Choice of Languages and Services
- Developers need a certain level of flexibility and control over how they develop.
- They need to be allowed to choose the best language and service for the job.
- What you choose to make available to individual teams and developers is configurable on a per team basis.

### The Open Source Ecosystem
- Open source is important, not just because it is free, but because anyone can get access to it.
- When an ecosystem is established, it becomes a safe choice for enterprise adoption because  
    you are aligning with a vibrant community pulling in the same direction.
- The Cloud Foundry ecosystem is not only backed by open source, it is backed by a foundation.
- This includes technical contributions to create software from a shared vision to support a common cause.
> Products and services that call themselves “Cloud Foundry” are
> making a commitment to users and to the broader ecosystem that
> they will offer a common experience across the vendors. For application
> developers, this means being able to deploy and manage
> applications in a consistent manner. For platform operations, this
> means that knowledge and skills are portable across the commercial
> products. It also means there are clearly defined integration
> points for the ecosystem to leverage to extend the platform’s core
> functionality.

### Cloud Foundry Constructs
- The Cloud Foundry platform offers:
    - Services as a higher level of abstraction above infrastructure:  
        Cloud Foundry provides a self-service mechanism for the ondemand deployment of applications bound to an array of  
        provisioned middleware services.  
        This benefit removes the management overhead of both the middleware and infrastructure layer from the developer,  
        significantly reducing the “development to deployment” time.
    - Containers: Cloud Foundry supports the use of container images such as Docker as a first-class citizen.  
        It also supports running applications artifacts deployed “as is,” containerizing them on the user’s behalf.  
        This flexibility allows companies already established with Docker to use their existing assets.  
        Containerizing applications on the user’s behalf offers additional productivity and operational benefits because  
        the resulting container image is built from known and vetted platform components,  
        leaving only the application source code to require vulnerability scanning.
    - Agile and automation: Cloud Foundry can be leveraged as part of a CI/CD pipeline to provision  
        environments and services on demand as the application moves through the pipeline to a production-ready state.  
        This helps satisfy the key Agile requirement of getting code into the hands of end users when required.
    - A cultural shift to DevOps: Cross-cutting concerns is a wellunderstood concept by developers.  
        Cloud Foundry is ideally accompanied by a cultural shift to DevOps.
    - Microservices support: Cloud Foundry supports microservices through providing mechanisms for integrating and coordinating loosely coupled services.  
        In order to realize the benefits of microservices, a platform is required to provide additional supporting capabilities,  
        such as built-in resilience and application authentication.
    - Cloud-native application support: Cloud Foundry provides a contract for applications to be developed against.  
        This contract makes doing the right thing the easy thing and will result in better application performance, management, and resilience.

P. 30
