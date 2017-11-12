## Introduction

DoneBot is a productivity coach and personal assistant bot. 

The purpose of DoneBot is to guide you through the process of getting things done: to keep you focused on the task at hand, to notice if you are stuck and ask the right questions to move things forward.

DoneBot also keeps an account of all the things you need to get done, just like any personal task manager. But where typical task managers are focused on the static _structure_ of getting things done, Donebot is focused on the _process_ of getting things done, and only keeps as much structure as is necessary to move your goals forward.

## Who is this for

Donebot is for people with lots of responsibilities.

Donebot is for people who do not have clearly defined assignments.

Donebot is for people who make up their own work.

Donebot is for people whose time is mostly unstructured.

Donebot is for people with few calendar appointments.

Donebot is for people who make an effort to put together a well-organized todo list, but can never keep it up to date.

Donebot is for people who are pretty productive, but can't shake the feeling that some things in their life never get any attention.

Donebot is for people who read Getting Things Done, by David Allen.

## How to use Donebot

Donebot starts with doing a "mind sweep" - going through all areas of your life and noting all things that are on your mind into Inbox.



First, as you go through your day, you put down reminders


## Interacting with DoneBot

You talk with DoneBot by pressing one of the labeled buttons. Occasionally you get a free text prompt, these are marked with a pencil icon (✏️). Even then you can pick one of the buttons as well.

There is also a set of "slash" commands that work anytime, like "/inbox Remember the milk" adds an inbox item any time.

## Concepts

### Inbox

Inbox is a collection of anything that comes up. Ideas, thoughts, promises made to other people, etc. Inbox is not a TODO list, so inbox items are not required to follow any format - it's all just freeform text. But please keep to one item per message, for easier processing.

### Goal

"Goal" is a very broad term. A Donebot goal is some outcome that you will achieve in the next one to two weeks. A goal might have a due date, or not. 

#### Trashing goals

Some goals you will put down are impractical, or untimely, or just not important enough to be completed. Have courage to delete these goals. It's not a failure - it's good housekeeping.

### Action

An "action" is the next step towards a goal. Every goal has exactly one next action at all times. The next action must such that you can immediately do it right now, if you have the time (and the context).

Actions are considered extremely expendable: you can always throw away the current next action and substitute a more appropriate one. It is expected that real world conditions change, and actions change with them. 

- If you discover the next action too complex,  write down the first primitive sub-action. 
- If the next action requires some preparation, put down the preparation step as the next action instead.

#### Planning

Donebot is not a tool for planning projects. I believe that, in most cases, only recording the next action works better in the real world. But for those goals that need explicit planning, you can set up a plan in an external program. 

Also, often, planning happens on a higher level, and lower-level items of the plan become Donebot goals.

### Context

Actions often have some physical prerequisites: location or access to people or tools, etc. So at any moment, some actions are physically impossible to do. To record that, you assign the action to a context. A context might be "home" or "office" or "computer" or "city". Donebot uses the context when picking the next action for you to do.

- A mindset could also be a context, such as "brainstorm" or "research".

### Waiting

A goal can become waiting if the next action is either delegated to another person, or scheduled to happpen at a certain time (like an appointment). In both these cases, you declare the goal as waiting. Waiting goals are reviewed regularly to see if you are ready to proceed.

## DoneBot conversations

There are several conversation flows that Donebot will take you through.

### Default mode - Collection

Collection is on all the time when you're not in the middle of a Donebot session. In Collection, any message you send to Donebot, gets recorded as an input item. This way you can quickly jot down notes anytime.

Then, when you have a block of time you want to spend productively (even if you're taking a bathroom break!), engage Donebot's work mode.

### Work mode

Work mode puts you through all stages of the workflow, from inbox item to goal completed.

### Inbox processing mode

If you have some items in your inbox, Donebot will prompt you to process them. 

"Processing" an inbox item involves either turning it into a goal and defining the next action towards it, or making use of the item then and there, and trashing.

### Doing mode

One of the core beliefs of Donebot is it's easier to get tough work done if you don't get to choose which action to do next. Thus, Donebot will choose actions for you.

#### Priority

Donebot deliberately does not let you prioritise one goal over another, apart from setting due dates. This is to avoid creating a stagnant pool of unimportant, eternally skipped, never finished goals. Instead, you cycle through all goals, and advance all of them. Goals that are not important enough to be worked on should be trashed.

### Waiting mode

When a goal needs some external action to happen, like when you delegate some action to another person, it enters waiting mode.

Waiting goals are checked daily, to see if the waiting is over. Due waiting goals are checked hourly.

### Review mode

Goals are reviewed weekly to make sure they are still relevant. When a goal is suggested for review, you can change its statement, or due date.

## Things Donebot does not do

### Checklists

Checklists are lists like "to read" or "to do before trip" or a shopping list. There are many nice apps to keep checklists, and I decided not to focus them, at least for now.

You can use a checklist _with_ Donebot - for example, checklist items might become Goals, or the entire checklist might be a Goal and you would pick the next action from the checklist.

### Calendar

Currently Donebot makes no attempt to replace your calendar app. 
Donebot is for work that is not scheduled. 

For actions that are scheduled, put them on the calendar, and set the goal as waiting in Donebot.

### Routine

Some work is repeated over and over, every day, or worse, _not_ every day. Sometimes it's just one action, sometimes an entire project. Donebot, for now, won't help you with organizing routine. Again, Donebot is for work that is not scheduled.

But, I recognize that routine work is in many ways similar to the work you'll put into Donebot, so supporting that is one of my future plans.

### Tickler

A "Tickler" system lets you set aside some inbox item to be reviewed at a later, predefined date. Like if you come across a cool Christmas present idea in May and defer it to September.

I use calendar reminders for my tickler file, but this is one of the feature I plan to add to Donebot.

### Reference

You need a system to keep reference material, like a notes system. Donebot isn't keeping track of your notes, or any project planning material.

Some inbox items you will put into donebox aren't actionable, but need to be put into your reference system. In such a case, file the inbox item manually, then trash it from Donebox.
