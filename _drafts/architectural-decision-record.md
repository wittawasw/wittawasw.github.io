---
layout: post
title: เก็บข้อมูลการตัดสินใจเชิงเทคนิคอย่างเป็นระบบด้วย ADR - Architectural Decision Record
---

### Definition and purpose of ADRs

An Architectural Decision Record (ADR) is a document that records and explains a significant architectural decision made during the development of a software system. The purpose of an ADR is to provide a clear, concise, and accessible record of the decision, the reasons behind it, and any alternatives that were considered.

An analogy that might help to illustrate the purpose of an ADR is to think of it as a "flight plan" for a software system. Just as a flight plan outlines the route and details of a journey by aircraft, an ADR outlines the key architectural choices and considerations that shape the development and evolution of a software system. Like a flight plan, an ADR serves as a reference and record of past decisions, and helps to guide future decision-making and planning. Overall, the purpose of an ADR is to ensure that the architectural choices made during the development of a software system are well-informed, transparent, and consistent with the overall goals and requirements of the system.

### Scenario
- Choosing a technology stack: When developing a new software system, there are often many different technology options to choose from. An ADR can be used to document the decision-making process and rationale behind the selection of a particular technology stack.
- Designing a system architecture: The overall architecture of a software system can have a significant impact on its performance, scalability, and maintainability. An ADR can be used to document and communicate the key design choices and trade-offs made during the architectural design process.
- Making changes to an existing system: As a software system evolves over time, it may be necessary to make changes to its architecture or design. An ADR can be used to document the reasons for these changes and the impact they are expected to have on the system.
- Resolving technical disputes: When there are differing opinions or approaches to a technical issue, an ADR can be used to facilitate a discussion and reach a consensus on the best course of action.
- Communicating with stakeholders: ADRs can be a useful tool for communicating technical decisions to stakeholders such as project managers, business analysts, and customers. By providing a clear and concise record of the decision and the reasoning behind it, ADRs can help to ensure that all stakeholders are informed and aligned on key technical issues.


## The benefits of using ADRs

### Improved communication and documentation of architectural decisions
In a software development project, it is important to have clear and consistent communication about the key technical choices and decisions that are being made. ADRs provide a centralized, documented record of these decisions, which can be used to communicate the rationale and implications of each decision to team members and stakeholders.

Having a documented record of architectural decisions can also be useful for future reference. As a software system evolves over time, it can be helpful to have a record of past decisions and the reasoning behind them. This can help to inform future decision-making and ensure that the system remains aligned with its original goals and requirements.

Overall, the use of ADRs can improve communication and documentation of architectural decisions by providing a clear, accessible, and consistent record of these decisions, which can be used to inform and guide future development efforts.

### Enhanced accountability and traceability of decisions
By documenting and recording architectural decisions in an ADR, it becomes easier to track and understand the decisions that have been made, and the impact they have had on the development of a software system. This can be especially useful when working with large, complex systems, where the consequences of a decision may not be immediately apparent.

In addition to improving traceability, ADRs can also enhance accountability by providing a clear record of who made a particular decision, and the reasoning behind it. This can help to ensure that decisions are made in a transparent and responsible manner, and that team members are held accountable for the choices they make.

Overall, the use of ADRs can improve the accountability and traceability of decisions by providing a documented record of the decisions that have been made, and the reasoning behind them. This can help to ensure that decisions are made in a responsible and transparent manner, and that the consequences of these decisions are well understood.

### Facilitates collaboration and consensus-building among team members
One of the key purposes of an ADR is to document and communicate the rationale and implications of an architectural decision. By providing a clear and accessible record of the decision, ADRs can help to ensure that all team members are informed and aligned on key technical issues. This can facilitate collaboration and coordination among team members, and help to avoid misunderstandings or conflicting approaches to problem-solving.

In addition, the use of ADRs can facilitate consensus-building by providing a structured process for discussing and resolving technical issues. By documenting and reviewing ADRs, team members can have a dialogue about the pros and cons of different approaches, and reach a consensus on the best course of action. This can help to ensure that decisions are made in a collaborative and inclusive manner, and that all team members feel that their perspectives and concerns have been taken into account.

Overall, the use of ADRs can facilitate collaboration and consensus-building among team members by providing a clear, documented record of technical decisions, and a structured process for discussing and resolving technical issues.

## The drawbacks of using ADRs

### Takes time and effort to create and maintain ADRs
Creating an ADR requires effort to document and communicate the details of a particular architectural decision. This includes identifying the problem or opportunity being addressed, outlining the alternatives that were considered, and explaining the reasoning behind the chosen solution. This process can be time-consuming, especially if the decision being documented is complex or has significant implications for the software system.

In addition to the time and effort required to create ADRs, there is also the ongoing effort required to maintain them. This may include updating the ADR to reflect changes to the system, or to correct any errors or inaccuracies that are discovered. Maintaining ADRs can be particularly important when working with large, complex systems, where the consequences of a decision may not be immediately apparent.

Overall, while the use of ADRs can provide many benefits, it is important to be aware that they also require time and effort to create and maintain. It is important to find the right balance and determine the appropriate level of effort to invest in ADRs, based on the needs and goals of the software development project.

### Requires discipline to consistently use the ADR process
Implementing ADRs effectively requires that team members consistently follow the process for documenting and reviewing architectural decisions. This includes identifying the decisions that should be documented in an ADR, creating the ADR in a timely manner, and reviewing and approving the ADR before it is finalized.

Maintaining discipline in the use of ADRs can be challenging, especially in fast-paced or high-pressure environments where the focus is on meeting deadlines and delivering features. It is important to establish clear guidelines and expectations for the use of ADRs, and to ensure that team members have the necessary support and resources to follow the ADR process consistently.

Overall, while the use of ADRs can provide many benefits, it is important to be aware that they also require discipline to consistently use the ADR process. It is important to establish clear guidelines and expectations for the use of ADRs, and to ensure that team members have the necessary support and resources to follow the ADR process consistently.

## A general approach to implementing ADRs in an organization

### Establishing guidelines for creating and reviewing ADRs
When implementing Architectural Decision Records (ADRs) in an organization, it is important to establish guidelines for creating and reviewing ADRs. These guidelines can help to ensure that ADRs are used consistently and effectively, and that the benefits of using ADRs are realized.

Some possible guidelines for creating and reviewing ADRs might include:

- Identifying the types of decisions that should be documented in an ADR
- Establishing a process for creating and reviewing ADRs, including who is responsible for creating and reviewing ADRs, and how ADRs should be reviewed and approved
- Setting standards for the content and format of ADRs, including the level of detail and documentation required for each ADR
- Defining the process for updating and maintaining ADRs, including how and when to correct errors or make changes to an ADR

Overall, establishing guidelines for creating and reviewing ADRs can help to ensure that the ADR process is implemented consistently and effectively, and that the benefits of using ADRs are realized.

### Identifying the appropriate level of detail and documentation for each ADR
When implementing Architectural Decision Records (ADRs), it is important to identify the appropriate level of detail and documentation for each ADR. This can help to ensure that ADRs are effective in documenting and communicating the key decisions and considerations related to a software system's architecture, while also being efficient and manageable to create and maintain.

One approach to determining the appropriate level of detail and documentation for an ADR is to consider the scope and significance of the decision being documented. For example, a decision with wide-ranging implications for the system may require more detailed documentation, while a more narrow or straightforward decision may require less.

It may also be helpful to consider the audience for the ADR when determining the appropriate level of detail and documentation. For example, an ADR intended for a technical audience may require more technical detail and explanation, while an ADR intended for a non-technical audience may need to be more high-level and accessible.

Overall, it is important to strike a balance between providing sufficient detail and documentation to effectively communicate the decision and its implications, while also keeping the ADR concise and manageable. Identifying the appropriate level of detail and documentation for each ADR can help to ensure that ADRs are effective and efficient in documenting and communicating architectural decisions.

### Ensuring that the ADR process is integrated into the overall software development process
When implementing Architectural Decision Records (ADRs), it is important to ensure that the ADR process is integrated into the overall software development process. This can help to ensure that ADRs are used consistently and effectively, and that the benefits of using ADRs are realized.

There are a few key ways in which the ADR process can be integrated into the overall software development process:

- Make the use of ADRs a standard part of the development process: By making the creation and review of ADRs a standard part of the development process, team members can be more likely to consistently follow the ADR process and realize the benefits of using ADRs.
- Integrate the ADR process into existing tools and processes: By integrating the ADR process into existing tools and processes, such as issue tracking systems or project management software, team members can more easily create and review ADRs as part of their normal workflow.
- Provide training and support for team members: To ensure that team members are able to effectively use ADRs, it may be helpful to provide training and support on the ADR process, including how to create and review ADRs, and how to use ADRs to facilitate collaboration and consensus-building.

Overall, by ensuring that the ADR process is integrated into the overall software development process, team members can be more likely to consistently use ADRs and realize the benefits of using them.

## Conclusion

ADRs can be a useful tool for improving communication, accountability, and collaboration in software development, but they also require effort to create and maintain. It is important to find the right balance and approach for using ADRs in an organization.
