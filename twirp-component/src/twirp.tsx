import React, { useContext, useState } from "react";

export { TwirpJSONClient } from "./twirp.json";

export interface TwirpClient<Req, Res> {
  request(method: string, variables: Partial<Req>, options: any): Promise<Res>;
}

interface TwirpClientContext<Req, Res> {
  client?: TwirpClient<Req, Res>;
}

export const TwirpContext = React.createContext<TwirpClientContext<any, any>>(
  {}
);

type TwirpMethod<Req, Res> = (r: Req, t: TwirpClient<Req, Res>) => Promise<Res>;
type Request<Req, Res> = (r: Req) => Promise<Res>;

export function useTwirp<Req, Res>(
  method: TwirpMethod<Req, Res>,
  req: Req
): Res;
export function useTwirp<Req, Res>(
  method: TwirpMethod<Req, Res>
): Request<Req, Res>;
export function useTwirp<Req, Res>(
  method: TwirpMethod<Req, Res>,
  req?: Req
): Res | Request<Req, Res> {
  const { client } = useContext(TwirpContext);
  const [res, setRes] = useState<Res | undefined>(undefined);

  if (typeof client == "undefined") {
    throw new Error(
      "twirp client missing. did you forget the <TwirpContext.Provider>?"
    );
  }

  if (typeof req == "undefined") {
    return (req: Req) =>
      method(req, client).then((res: Res) => {
        setRes(res);
        return res;
      });
  }

  if (typeof res == "undefined") {
    throw method(req, client).then((res: Res) => {
      setRes(res);
      return res;
    });
  }

  return res;
}
