## Assignment

### Implementing Target Systems with Interfaces

#### Objective

This assignment will help you understand how to define and work with interfaces in Go by simulating different types of "Target Systems" for sending notifications or updates to different external systems such as an Email system, SMS system, and Push notification system. The main goal is to abstract the notification system by using a common interface.

#### Problem Statement

You are building a system that sends notifications to various external target systems (e.g., email, SMS, or push notification). The requirement is to design the system in such a way that adding a new target (e.g., a Slack notification or a webhook) in the future should be easy, without changing the core notification sending logic.

To achieve this, you need to define a `Target` interface that will be implemented by different target types (Email, SMS, and Push). Each target system will handle the message sending in its own way, but the main program will interact with them through the `Target` interface.

#### Requirements

1. Define the <code>Target</code> Interface</strong>:
   - The <code>Target</code> interface should have one method:
     - <code>Send(message string) error</code>: This will send the message to the target system (e.g., an email, SMS, or push notification).
2. <strong>Implement at least three types</strong> that satisfy the <code>Target</code> interface:
   - <strong>EmailTarget</strong>: Simulates sending an email by printing the message to the console.
   - <strong>SMSTarget</strong>: Simulates sending an SMS message by printing the message to the console.
   - <strong>PushNotificationTarget</strong>: Simulates sending a push notification by printing the message to the console.
   - <strong>(Optional)</strong> You can also implement a <strong>SlackTarget</strong> to simulate sending a Slack message or a <strong>WebhookTarget</strong> to simulate sending a webhook notification.
3. <strong>Add a Notification Manager</strong>:
   - Implement a <code>NotificationManager</code> struct that holds a list of <code>Target</code>s and sends the same message to all targets.
4. <strong>Switch Between Different Target Systems</strong>:
   - Use command-line flags or input arguments to specify which target systems should be used (e.g., email, SMS, push, or multiple).
