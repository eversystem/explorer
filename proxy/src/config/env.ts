import { resolve } from "path";
import { config } from "dotenv";

config({ path: resolve(__dirname, "../../.env") });

const env: { [key: string]: string } = {
  BASICNAME: process.env.BASICNAME || "",
  BASICPASS: process.env.BASICPASS || "",
  JWTSEC: process.env.JWTSEC || ""
};

Object.keys(env).forEach((key: string) => {
  if (!env[key]) {
    throw new Error(`env: ${key} is not defined!!`);
  }
});

export default env;
