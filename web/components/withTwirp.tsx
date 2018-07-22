import React from "react";
import { NextContext } from "next";

import {
  renderState,
  TwirpJSONClient as TwirpClient,
  TwirpProvider,
  InMemoryCache
} from "../rpc/twirp";

const prefix = "http://localhost:4000/twirp/";

type Props = {
  ssrCache?: string;
} & { children?: React.ReactNode };

const withTwirp = (Component: React.ComponentClass | React.SFC) =>
  class extends React.Component<Props> {
    static async getInitialProps({ req }: NextContext) {
      const cache = new InMemoryCache();
      // only render state in ssr. let client side navigation
      // show loading states instead
      if (req) {
        const client = new TwirpClient(prefix);
        try {
          console.log(" -- APP RENDER STATE START");
          await renderState(client, cache, <Component />);
          console.log(" -- APP RENDER STATE END");
        } catch (err) {
          console.error("renderState err", err);
        }
      }
      return { ssrCache: await cache.dump() };
    }

    twirp = {
      client: new TwirpClient(prefix),
      cache: new InMemoryCache().load(this.props.ssrCache)
    };

    render() {
      return (
        <TwirpProvider value={this.twirp}>
          <Component {...this.props} />
        </TwirpProvider>
      );
    }
  };

export default withTwirp;