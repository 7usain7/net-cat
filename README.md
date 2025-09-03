# net-cat

## Summary

This project is a basic implementation of a NetCat-like tool built with Go. It allows multiple clients to connect to a server using TCP, send and receive messages in real-time, and acts like a simple group chat.

## Features

- TCP server that handles up to 10 clients at once
- Each client has to enter a name when joining
- All messages are time-stamped and labeled with the sender name
- Clients receive all past messages when they join
- Every user connected is notified when a new user joins or leaves
- Messages are shared in real-time to all connected clients

## Message Format

- All messages follow this format:

[YYYY-MM-DD HH:MM:SS][username]: message

## How to Run

`go run main.go 8989`

## How to connect client

`nc localhost 8989`

## Team Members

- Hussain Abdulrasool #habdulras

- Ali Madan #alimadan
