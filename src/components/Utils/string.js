export default function truncate(input, seperator) {
  // Invent a Number if seperator is not passed in
  // eslint-disable-next-line no-param-reassign
  seperator = seperator !== undefined ? seperator : 10;
  if (input.length > seperator) {
    return `${input.substring(0, seperator)} ...`;
  }
  return input;
}
