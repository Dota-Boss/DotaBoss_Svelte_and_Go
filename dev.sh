#!/bin/bash

# Start Go backend
echo "Starting backend..."
cd backend
go run main.go &
BACKEND_PID=$!
cd ..

# Start Svelte frontend
echo "Starting frontend..."
cd frontend
npm run dev &
FRONTEND_PID=$!
cd ..

# Wait and trap exit
trap "kill $BACKEND_PID $FRONTEND_PID" SIGINT
wait
