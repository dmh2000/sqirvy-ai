#!/bin/bash

# USE REASONER
../../bin/sqirvy-query -m deepseek-reasoner <python-web-prompt.md >peasoner.md

# USE CHAT
../../bin/sqirvy-query -m deepseek-chat <python-web-prompt.md >chat.md