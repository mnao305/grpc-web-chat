import { MessengerClient } from './MessengerServiceClientPb'

export type MessengerClientAttached = {
  client: MessengerClient;
};

export const client = new MessengerClient('http://0.0.0.0:8081')
