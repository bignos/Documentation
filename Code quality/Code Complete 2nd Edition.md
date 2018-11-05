Code Complete 2nd Edition
=========================

--------------------------------------------------------------------------------

## Welcome to software construction

1. What Is Software Construction?

### Software developement activities:
- Problem definition
- Requirements development
- Construction planning
- Software architecture, or high-level design
- Detailed design
- Coding and debugging
- Unit testing
- Integration testing
- Integration
- System testing
- Corrective maintenance

### Construction activities:
_Construction = Programming_

- Verifying that the groundwork has been laid so that construction can proceed successfully
- Determining how your code will be tested
- Designing and writing classes and routines
- Creating and naming variables and named constants
- Selecting control structures and organizing blocks of statements
- Unit testing, integration testing, and debugging your own code
- Reviewing other team members’ low-level designs and code and having them review yours
- Polishing code by carefully formatting and commenting it
- Integrating software components that were created separately
- Tuning code to make it smaller and faster

2. Why Is Software Construction Important?

- Construction is a large part of software development
- Construction is the central activity in software development
- With a focus on construction, the individual programmer’s productivity can improve enormously
- Construction’s product, the source code, is often the only accurate description of the software
- Construction is the only activity that’s guaranteed to be done

3. Key Points

- Software construction the central activity in software development
- Construction is the only activity that’s guaranteed to happen on every project.
- The main activities in construction are detailed design, coding, debugging, and developer testing.
- Other common terms for construction are “coding and debugging” and “programming.”
- The quality of the construction substantially affects the quality of the software.

--------------------------------------------------------------------------------

## Metaphors for a Richer Understanding of Software Development

1. The Importance of Metaphors
_Metaphors = Models_

2. How to Use Software Metaphors

- Use them to give you insight into your programming problems and processes.
- Use them to help you think about your programming activities and to help you imagine better ways of doing things.

3. Common Software Metaphors

- Software Penmanship: Writing Code
- Software Farming: Growing a System
- Software Oyster Farming: System Accretion
    - _Accretion = incremental, iterative, adaptive, evolutionary_
    - Incremental designing, building, and testing are some of the most powerful software development concepts available.
- Software Construction: Building Software
    - The penalty for a mistake on a simple structure is only a little time and maybe some embarrassment.
    - More complicated structures require more careful planning
    - Likewise, in software you might generally use flexible, lightweight development approaches
    - But sometimes rigid, heavyweight approaches are required to achieve safety goals and other goals.
- Applying Software Techniques: The Intellectual Toolbox
    - The more you learn about programming, the more you fill your mental toolbox 
    -   with analytical tools and the knowledge of when to use them and how to use them correctly.
    - The toolbox metaphor helps to keep all the methods, techniques, and tips in perspective

4. Key Points

- Metaphors are heuristics, not algorithms. As such, they tend to be a little sloppy.
- Metaphors help you understand the software-development process by relating it to other activities you already know about.
- Some metaphors are better than others.
- Treating software construction as similar to building construction 
    suggests that careful preparation is needed and illuminates the difference between large and small projects.
- Thinking of software-development practices as tools in an intellectual toolbox 
    suggests further that every programmer has many tools and that no single tool is right for every job.
- Choosing the right tool for each problem is one key to being an effective programmer.

--------------------------------------------------------------------------------

## Measure Twice, Cut Once: Upstream Prerequisites

1. Importance of Prerequisites

- Why proper preparation is important and tells you how to determine whether you’re really ready to begin construction.

### Do Prerequisites Apply to Modern Software Pro-jects?
- projects will run best if appropriate preparation activities are done before construction begins in earnest.

### Causes of Incomplete Preparation

### Utterly Compelling and Foolproof Argument for Doing Prerequisites Before Construction
- Part of your job as a technical employee is to educate the nontechnical people around you about the development process.

### Appeal to Logic
- It’s also important to think about how to build the system before you begin to build it.
- You don’t want to spend a lot of time and money going down blind alleys 
    when there’s no need to, especially when that increases costs.

### Appeal to Analogy
- If you are planning a highly iterative project:
    - You will need to identify the critical requirements and architectural elements 
        that apply to each piece you’re constructing before you begin construction.

### Appeal to Data
- Studies over the last 25 years have proven conclusively that it pays to do things right the first time.
- Unnecessary changes are expensive.
- Dozens of companies have found that simply focusing on correcting defects earlier 
    rather than later in a project can cut development costs and schedules by factors of two or more.
- This is a healthy incentive to fix your problems as early as you can.

2. Determine the Kind of Software You’re Working On

- Iterative approaches tend to reduce the impact of inadequate upstream work, but they don’t eliminate it.
- One realistic approach is to plan to specify about 80 percent of the requirements
    up front, allocate time for additional requirements to be specified later, and then
    practice systematic change control to accept only the most valuable new re-
    quirements as the project progresses.
- Another alternative is to specify only the most important 20 percent of the re-
    quirements up front and plan to develop the rest of the software in small incre-
    ments, specifying additional requirements and designs as you go.

3. Problem-Definition Prerequisite

- Problem definition comes before detailed requirements work, which is a more indepth investigation of the problem.
- The problem definition should be in user language, and the problem should be described from a user’s point of view. 
- Without a good problem definition,
    you might put effort into solving the wrong problem.
    Be sure you know what you’re aiming at before you shoot.

4. Requirements Prerequisite
_requirements = specification_

### Why Have Official Requirements ?
- If the requirements are explicit, the user can review them and agree to them.
- If they’re not, the programmer usually ends up making requirements decisions during programming.
- If you have a disagreement with an other programmer about what the program is supposed to do, you can resolve it
    by looking at the written requirements.
- Without good requirements, you can have the right general problem but miss the
    mark on specific aspects of the problem.

### The Myth of Stable Requirements

### Handling Requirements Changes During Construction
- Use the requirements checklist at the end of the section to assess the quality of your requirements
- Make sure everyone knows the cost of requirements changes
- Set up a change-control procedure
- Use development approaches that accommodate changes
    - The key is using short development cycles so that you can respond to your users quickly.
- Dump the project

#### Checklist: Requirements
- Specific Functional Requirements:
    - Are all the inputs to the system specified,
        including their source, accuracy, range of values, and frequency ?
    - Are all the outputs from the system specified,
        including their destination, accuracy, range of values, frequency, and format ?
    - Are all output formats specified for web pages, reports, and so on ?
    - Are all the external hardware and software interfaces specified ?
    - Are all the external communication interfaces specified,
        including handshaking, error-checking, and communication protocols ?
    - Are all the tasks the user wants to perform specified ?
    - Is the data used in each task and the data resulting from each task specified ?

- Specific Non-Functional (Quality) Requirements:
    - Is the expected response time, from the user’s point of view, specified for all necessary operations ?
    - Are other timing considerations specified, such as processing time, datatransfer rate, and system throughput ?
    - Is the level of security specified ?
    - Is the reliability specified, including the consequences of software failure,
        the vital information that needs to be protected from failure, and the strategy
        for error detection and recovery ?
    - Is maximum memory specified ?
    - Is the maximum storage specified ?
    - Is the maintainability of the system specified, including its ability to adapt to
        changes in specific functionality, changes in the operating environment, and
        changes in its interfaces with other software ?
    - Is the definition of success included? Of failure ?

- Requirements Quality:
    - Are the requirements written in the user’s language ? Do the users think so ?
    - Does each requirement avoid conflicts with other requirements ?
    - Are acceptable trade-offs between competing attributes specified
        for example, between robustness and correctness ?
    - Do the requirements avoid specifying the design ?
    - Are the requirements at a fairly consistent level of detail ?
        Should any requirement be specified in more detail ?
        Should any requirement be specified in less detail ?
    - Are the requirements clear enough to be turned over to an independent group for construction and still be understood ?
    - Is each item relevant to the problem and its solution ?
        Can each item be traced to its origin in the problem environment ?
    - Is each requirement testable ?
        Will it be possible for independent testing to determine whether each requirement has been satisfied ?
    - Are all possible changes to the requirements specified, including the likelihood of each change ?

- Requirements Completeness:
    - Where information isn’t available before development begins, are the areas of incompleteness specified ?
    - Are the requirements complete in the sense that if the product satisfies every requirement, it will be acceptable ?
    - Are you comfortable with all the requirements ?
        Have you eliminated requirements that are impossible to implement and included just to appease
        your customer or your boss ?

5. Architecture Prerequisite

- Good architecture makes construction easy. Bad architecture makes construction almost impossible.
- Without good software architecture, you may have the right problem but the wrong solution.
    It may be impossible to have successful construction.

### Typical Architectural Components
- Program Organization
    - One review of design practices found that the design rationale is at least as important
        for maintenance as the design itself
    - Every feature listed in the requirements should be covered by at least one building block.
        If a function is claimed by two or more building blocks, their claims should cooperate, not conflict.
    - By minimizing what each building block knows about each other building block,
        you localize information about the design into single building blocks.
    - The architecture should describe which other building blocks the building block can
        use directly, which it can use indirectly, and which it shouldn’t use at all.
- Major Classes
    - specify the 20 percent of the classes that make up 80 percent of the systems behavior
- Data Design
    - The architecture should describe the major files and table designs to be used.
    - Data should normally be accessed directly by only one subsystem or class.
    - The architecture should specify the high-level organization and contents of any databases used.
- Business Rules
    - If the architecture depends on specific business rules, it should identify them and
        describe the impact the rules have on the system’s design.
- User Interface Design
    - Careful architecture of the user interface
        makes the difference between a wellliked program and one that’s never used.
    - The architecture should be modularized so that a new user interface 
        can be substituted without affecting the business rules and output parts of the program.
- Input/Output
    - Input/output is another area that deserves attention in the architecture.
    - The architecture should specify a look-ahead, look-behind, or just-in-time reading scheme.
    - And it should describe the level at which I/O errors are detected: at the field, record, stream, or file level.
- Resource Management
    - The architecture should describe a plan for managing scarce resources such as database connections, threads, and handles.
    - The resource manager should be architected as carefully as any other part of the system.
- Security
    - The architecture should describe the approach to design-level and code-level security.
- Performance
    - Performance goals can include both speed and memory use.
- Scalability
    - The architecture should describe how the system will address growth in
        number of users, number of servers, number of network nodes, database size, transaction volume.
    - If the system is not expected to grow and scalability is not an issue,
        the architecture should make that assumption explicit.
- Interoperability
    - If the system is expected to share data or resources with other software or hardware,
        the architecture should describe how that will be accomplished.
- Internationalization/Localization
    - Internationalization issues deserve attention in the architecture for an interactive system.
    - Translating the strings into foreign languages with minimal impact on the code and the user interface.
    - keep the strings in a class and reference them through the class interface, or store the strings in a resource file.
- Error Processing
    - A strategy for handling them consistently should be spelled out in the architecture.
    - It is best treated at the architectural level.
    - Here are some questions to consider:
        - Is error processing corrective or merely detective ?
             - In either case, it should notify the user that it detected an error.
        - Is error detection active or passive ?
        - How does the program propagate errors ?
        - What are the conventions for handling error messages ?
            - The architecture should establish conventions for error messages.
        - Inside the program, at what level are errors handled ?
        - What is the level of responsibility of each class for validating its input data ?
        - Do you want to use your environment’s built-in exception handling mechanism, or build your own ?
- Fault Tolerance
    - The architecture should also indicate the kind of fault tolerance expected.
    - Examples:
        - The system might back up and try again when it detects a fault.
            - it would back up to a point at which it knew everything was all right and continue from there.
        - The system might have auxiliary code to use if it detects a fault in the primary code.
        - The system might use a voting algorithm.
        - The system might replace the erroneous value with a phony value
            that it knows to have a benign effect on the rest of the system.
- Architectural Feasibility
    - The architecture should demonstrate that the system is technically feasible.
    - These risks should be resolved before full-scale construction begins.
- Overengineering
    - The architecture should clearly indicate whether programmers should err on the side of overengineering 
        or on the side of doing the simplest thing that works.
    - By setting expectations explicitly in the architecture,
        you can avoid the phenomenon in which some classes are exceptionally robust and others are barely adequate.
- Buy-vs.-Build Decisions
    - The most radical solution to building software is not to build it at all—to use external software instead.
- Reuse Decisions
    - If the plan calls for using pre-existing software,
        the architecture should explain how the reused software will be made to conform to the other architectural goals
- Change Strategy
    - One of the major challenges facing a software architect is
        making the architecture flexible enough to accommodate likely changes.
    - The architecture should clearly describe a strategy for handling changes.
- General Architectural Quality
    - A good architecture specification is characterized by discussions of the classes in the system,
        of the information that’s hidden in each class,
        and of the rationales for including and excluding all possible design alternatives.
    - The essential problem with large systems is maintaining their conceptual integrity
    - A good architecture should fit the problem.
    - When you look at the architecture, you should be pleased by how natural and easy the solution seems.
    - The architecture’s objectives should be clearly stated.
    - The architecture should describe the motivations for all major decisions.
    - Good software architecture is largely machine and language independent.
    - The architecture should address all requirements without over-specifying the system.
    - The architecture should explicitly identify risky areas.
    - It shouldn’t contain anything just to please the boss.
    - It shouldn’t contain anything that’s hard for you to understand.

#### Checklist: Architecture
- Specific Architectural Topics:
    - Is the overall organization of the program clear, including a good architectural overview and justification ?
    - Are major building blocks well defined, including their areas of responsibility and their interfaces to other building blocks ?
    - Are all the functions listed in the requirements covered sensibly, by neither too many nor too few building blocks ?
    - Are the most critical classes described and justified ?
    - Is the data design described and justified ?
    - Is the database organization and content specified ?
    - Are all key business rules identified and their impact on the system described ?
    - Is a strategy for the user interface design described ?
    - Is the user interface modularized so that changes in it won’t affect the rest of the program ?
    - Is a strategy for handling I/O described and justified ?
    - Are resource-use estimates and a strategy for resource management described and justified ?
    - Are the architecture’s security requirements described ?
    - Does the architecture set space and speed budgets for each class, subsystem, or functionality area ?
    - Does the architecture describe how scalability will be achieved ?
    - Does the architecture address interoperability ?
    - Is a strategy for internationalization/localization described ?
    - Is a coherent error-handling strategy provided ?
    - Is the approach to fault tolerance defined ?
    - Has technical feasibility of all parts of the system been established ?
    - Is an approach to overengineering specified ?
    - Are necessary buy-vs.-build decisions included ?

- General Architectural Quality:
    - Does the architecture account for all the requirements ?
    - Is any part over- or under-architected ?
        - Are expectations in this area set out explicitly ?
    - Does the whole architecture hang together conceptually ?
    - Is the top-level design independent of the machine and language that will be used to implement it ?
    - Are the motivations for all major decisions provided ?
    - Are you, as a programmer who will implement the system, comfortable with the architecture ?

6. Amount of Time to Spend on Upstream Prerequisites
- Estimate the time for the rest of the project after you’ve finished the requirements.
- Ensure that the time you need to create a good architecture 
    won’t take away from the time you need for good work in other areas.
-  If necessary, plan the architecture work as a separate project too.

#### Checklist: Upstream Prerequisites
- Have you identified the kind of software project you’re working on and tailored your approach appropriately ?
- Are the requirements sufficiently well-defined and stable enough to begin construction ?
- Is the architecture sufficiently well defined to begin construction ?
- Have other risks unique to your particular project been addressed,
    such that construction is not exposed to more risk than necessary ?

7. Key Points
- The overarching goal of preparing for construction is risk reduction.
    Be sure your preparation activities are reducing risks, not increasing them.
- If you want to develop high-quality software, attention to quality must 
    be part of the software-development process from the beginning to the end.
    Attention to quality at the beginning has a greater influence on product quality than attention at the end.
- Part of a programmer’s job is to educate bosses and coworkers about the software-development process,
    including the importance of adequate preparation before programming begins.
- The kind of project you’re working significantly affects construction prerequisites
    many projects should be highly iterative, and some should be more sequential.
- If a good problem definition hasn’t been specified, you might be solving the wrong problem during construction.
- If a good requirements work hasn’t been done, you might have missed important details of the problem.
    Requirements changes cost 20 to 100 times as much in the stages following construction as they do earlier,
    so be sure the requirements are right before you start programming.
- If a good architectural design hasn’t been done, you might be solving the right problem the wrong way during construction. 
    The cost of architectural changes increases as more code is written for the wrong architecture,
    so be sure the architecture is right too.
- Understand what approach has been taken to the construction prerequisites on your project
    and choose your construction approach accordingly.

--------------------------------------------------------------------------------

## Key Construction Decisions

1. Choice of Programming Language
_By relieving the brain of all unnecessary work, a good notation sets it free to concentrate 
    on more advanced problems, and in effect increases the mental power of the race._

- Studies have shown that the programming-language choice affects productivity and code quality in several ways.
- Programmers are more productive using a familiar language than an unfamiliar one.
- Programmers working with high-level languages achieve better productivity and quality than those working with lower-level languages.
- Developers working in interpreted languages tend to be more productive than those working in compiled languages.
- Some languages are better at expressing programming concepts than others.

### Language Descriptions
- Ada
    Ada is used primarily in military, space, and avionics systems.
- Assembly Language
    Most programmers avoid it unless they’re pushing the limits in execution speed or code size.
- C
- C++
- C#
- Cobol
    The acronym Cobol stands for Common Business Oriented Language.
- Fortran
    Fortran is used mainly in scientific and engineering applications.
- Java
    Java is in widespread use for programming Web applications.
- JavaScript
     It is used primarily for adding simple functions and online applications to web pages.
- Perl
    The acronym Perl stands for Practical Extraction and Report Language.
- PHP
- Python
- SQL

2. Programming Conventions

- Before construction begins, spell out the programming conventions you’ll use.
    They’re at such a low level of detail that they’re nearly impossible to retrofit into software after it’s written.

3. Your Location on the Technology Wave

4. Selection of Major Construction Practices

- Part of preparing for construction is deciding which of the many available good practices you’ll emphasize.

#### Checklist: Major Construction Practices
- Coding:
    - Have you defined coding conventions for names, comments, and formatting ?
    - Have you defined specific coding practices that are implied by the architecture,
        such as how error conditions will be handled, how security will be addressed, and so on ?
    - Have you identified your location on the technology wave and adjusted your approach to match ?
        If necessary, have you identified how you will program into the language rather than being limited by programming in it ?

- Teamwork:
    - Have you defined an integration procedure, that is,
        have you defined the specific steps a programmer must go through before checking code into the master sources ?
    - Will programmers program in pairs, or individually, or some combination of the two ?

- Quality Assurance:
    - Will programmers write test cases for their code before writing the code itself ?
    - Will programmers write unit tests for the their code regardless of whether they write them first or last ?
    - Will programmers step through their code in the debugger before they check it in ?
    - Will programmers integration-test their code before they check it in ?
    - Will programmers review or inspect each other's code ?

- Tools:
    - Have you selected a revision control tool ?
    - Have you selected a language and language version or compiler version ?
    - Have you decided whether to allow use of non-standard language features ?
    - Have you identified and acquired other tools you’ll be using
        editor, refactoring tool, debugger, test framework, syntax checker, and so on ?

5. Key Points
- Every programming language has strengths and weaknesses.
    Be aware of the specific strengths and weaknesses of the language you’re using.
- Establish programming conventions before you begin programming.
    It’s nearly impossible to change code to match them later.
- More construction practices exist than you can use on any single project.
    Consciously choose the practices that are best suited to your project.
- Your position on the technology wave determines what approaches will be effective or even possible.
    Identify where you are on the technology wave, and adjust your plans and expectations accordingly.

--------------------------------------------------------------------------------

## Design in Construction


--------------------------------------------------------------------------------
