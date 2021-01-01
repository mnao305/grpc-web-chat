import { MessengerClient } from './MessengerServiceClientPb'

export type MessengerClientAttached = {
  client: MessengerClient;
};

export const client = new MessengerClient(`http://${location.hostname}:8081`)
