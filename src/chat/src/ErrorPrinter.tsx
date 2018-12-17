import React from 'react';

interface Props {
  err: any
}

export default function PrintError(props: Props) {
  const err = props.err;
  if (err["networkError"]["result"]["errors"] instanceof Array) {
    const errList = err["networkError"]["result"]["errors"];
    return (
      <div>
        Errors:
        <ul>
          {errList.map((err: any, i: Number) => {
            return <li key="err-{i}">{err.message}</li>
          })}
        </ul>
      </div>
    )
  }
  return <div>Error! ${err.message}</div>;
}

