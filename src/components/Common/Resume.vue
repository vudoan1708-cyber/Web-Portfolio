<template>
  <div id="resume">
    <h1>RESUME</h1> <br />
    <div id="embeded">
      <p>Open the CV file in another
        <a href="https://drive.google.com/file/d/1vfeW2XzxVGepOHKfwQwSPq_Xl32BFCaS/view?usp=sharing" target="_blank">window</a>
      </p> <br />
      <iframe :src="CV" width="100%" height="750px">
      </iframe>
    </div>
  </div>
</template>

<script>
// PDF
import CV from '@/assets/pdfs/CV.pdf';

export default {
  name: 'Resume',
  setup() {
    return {
      CV,
    };
  },
};
</script>

<style lang="scss" scoped>
#resume {
  display: block;
  position: relative;
  top: 15%;
  width: 80%;
  padding: 20px;
  color: snow;

  #embeded {
    p {
      text-align: right;
    }
  }
}
</style>
