import React, {useEffect, useState} from 'react';
import gql from 'graphql-tag';
import {useApolloClient, useMutation} from "react-apollo-hooks";

type Msg = {
    id: string
    text: string
}
type Room = {
    messages: Msg[]
}

function useRoom(channelName: string): [any, Boolean, Function] {
    const [room, setRoom] = useState<Room>({
        messages: [],
    });
    const client = useApolloClient();

    const [loading, setLoading] = useState(false);

    function handleQueryResult(result: any) {
      console.log(result)
        setLoading(false);
        setRoom(result.data.room);
    }

  useEffect(() => {

          if (!client) return;

      client.query({
        query: gql`
            query Room($channel: String!) {
                room(name: $channel) {
                    messages { id text createdBy createdAt }
                }
            }
        `,
          variables: {
              channel: channelName,
          },
      }).then(handleQueryResult);
    },
      [channelName, client]
  );

    return [room, loading, setRoom];
}

function useSubscription(channelName: string, setRoom: Function): void {
    const client = useApolloClient();

    function handleNewMessage(result: any) {
        if (!result.data) {
            return
        }
        setRoom((room: any) => {
            const newMessage = result.data.messageAdded;
            if (room.messages.find((msg: any) => msg.id === newMessage.id)) {
                return room
            }

            room.messages.push(newMessage);
            return room;
        });
    }

  useEffect(() => {
      if (!client) return;

    const subscription = client.subscribe({
      query: gql`
          subscription MoreMessages($channel: String!) {
              messageAdded(roomName:$channel) {
                  id
                  text
                  createdBy
              }
          }
      `,
        variables: {
            channel: channelName,
        },
    }).subscribe(handleNewMessage);

      return () => subscription.unsubscribe();
  }, [channelName]);
}

function Room(props: any) {
    const [room, loading, setRoom] = useRoom(props.channel);
    useSubscription(props.channel, setRoom);
    const [text, setText] = useState("");
    const onSubmit = useMutation(Mutation);

    // if (error) return <div>Error! ${error.message}</div>;
    if (loading) return <div>loading</div>;

    return <div>
        <div>
            {room.messages.map((msg: any) =>
                <div key={msg.id}>{msg.createdBy}: {msg.text}</div>
            )}
        </div>
        <input value={text} onChange={(e) => setText(e.target.value)} />
        <button onClick={() => onSubmit({
            variables: {
                text: text,
                channel: props.channel,
                name: props.name,
            }
        })}>send
        </button>
    </div>;
}

const Mutation = gql`
    mutation sendMessage($text: String!, $channel: String!, $name: String!) {
        post(text:$text, roomName:$channel, username:$name) { id }
    }
`;


export default Room
