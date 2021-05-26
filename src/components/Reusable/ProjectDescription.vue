<template>
  <div class="descriptions">
    <h2>{{ projectDetails.title }}</h2>
    <div class="details">
      <section>
        <h4>Description: </h4>
          <p>{{ projectDetails.description }} {{ projectDetails.orgName }}
          </p> <br />
      </section>

      <section>
        <h4>Project Type: </h4>
          <p>{{ projectDetails.projectType }}</p> <br />
      </section>

      <section>
        <h4>Role: </h4>
          <p>{{ projectDetails.role }}</p> <br />
      </section>

      <section v-if="projectDetails.programmingLangs.length > 0">
        <h4>Programming Languages: </h4>
          <p v-for="(lang, langKey) in projectDetails.programmingLangs"
            :key="langKey">{{ lang }}
          </p> <br />
      </section>

      <!-- <h4>Github Link: </h4>
        <p>{{ projectDetails.githubURL }}</p> <br />

      <h4>Project Link: </h4>
        <p>{{ projectDetails.liveURL }}</p> <br /> -->
    </div>
  </div>
</template>

<script>
import { reactive, ref, watch } from 'vue';

// Utils
import truncate from '@/components/Utils/string';

export default {
  name: 'ProjectDescription',
  props: {
    description: {
      type: String,
    },
    githubURL: {
      type: String,
    },
    liveURL: {
      type: [Object, String],
    },
    programmingLangs: {
      type: Array,
    },
    projectType: {
      type: String,
    },
    role: {
      type: String,
    },
    title: {
      type: String,
    },
    videos: {
      type: Array,
    },
    presentations: {
      type: Array,
    },
    orgName: {
      type: String,
    },
    orgExtraLink: {
      type: String,
    },
    mobile: {
      type: Boolean,
    },
  },
  setup(props) {
    // Props
    const projectDetails = reactive({
      description: props.description,
      githubURL: props.githubURL,
      liveURL: props.liveURL,
      programmingLangs: props.programmingLangs,
      projectType: props.projectType,
      role: props.role,
      title: props.title,
      orgName: props.orgName,
      orgExtraLink: props.orgExtraLink,
    });
    // Mobile Detection
    const isMobile = ref(props.mobile);
    // Truncate String To ... If In Mobile Mode
    if (isMobile.value) projectDetails.description = truncate(projectDetails.description, 50);

    watch(() => [
      props.description, props.githubURL, props.liveURL, props.programmingLangs,
      props.projectType, props.role, props.title, props.orgName, props.orgExtraLink, props.mobile,
    ],
    ([
      desc, githubURL, liveURL, programmingLangs,
      projectType, role, title, orgName, orgExtraLink, mobile,
    ]) => {
      projectDetails.description = desc;
      projectDetails.githubURL = githubURL;
      projectDetails.liveURL = liveURL;
      projectDetails.programmingLangs = programmingLangs;
      projectDetails.projectType = projectType;
      projectDetails.role = role;
      projectDetails.title = title;
      projectDetails.orgName = orgName;
      projectDetails.orgExtraLink = orgExtraLink;

      isMobile.value = mobile;
    });

    return {
      projectDetails,
    };
  },
};
</script>

<style lang="scss" scoped>
.descriptions {
  position: inherit;
  top: 0;
  left: 0;
  color: rgb(44, 44, 44);
  width: 100%;
  height: 100%;
  border-top-left-radius: 20px;
  border-bottom-left-radius: 20px;

  &:before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    border-top-left-radius: 20px;
    border-bottom-left-radius: 20px;
    box-shadow: inset 0 0 2000px rgba(248, 0, 0, 0.5);
    filter: blur(5px);
    z-index: -10;
  }

  h2 {
    margin: 20px;
  }
  .details {
    margin: 20px;
    text-align: left;
  }
}

/* Responsive Text */
@media screen and (max-width: 770px) {
  .descriptions {
    font-size: 15px;
  }
}
@media screen and (max-width: 600px) {
  .descriptions {
    font-size: 12px;
  }
}
@media screen and (max-width: 460px) {
  .descriptions {
    font-size: 9px;
  }
}
</style>
