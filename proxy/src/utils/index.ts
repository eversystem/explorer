import jwt, { SignOptions } from "jsonwebtoken";
import env from "../config/env";

const jwtSecret = env.JWTSEC;
const jwtOptions = {
  algorithm: "HS256",
  expiresIn: 3600
} as SignOptions;

export const generateJWT = () => jwt.sign({}, jwtSecret, jwtOptions);

export const veryfyJWT = (token: string): boolean => {
  let isValidJWT = false;
  jwt.verify(token, jwtSecret, (err: any, decoded: any) => {
    // eslint-disable-next-line no-empty
    if (err) {
    } else {
      isValidJWT = true;
    }
  });
  return isValidJWT;
};
