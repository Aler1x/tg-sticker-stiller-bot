# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-10-26

### Initial realease of bot

## Features

- Copy sticker packs
- Copy emoji packs
- List packs
- Delete packs
- Broadcast messages to all users (admin only)
- View bot statistics (admin only)

## Internal changes

- Use webhooks for production (Railway)
- Use long polling for local development
- Use SQLite database for persistence
- Use Telebotv4 framework for Telegram API
- Use Docker for deployment
