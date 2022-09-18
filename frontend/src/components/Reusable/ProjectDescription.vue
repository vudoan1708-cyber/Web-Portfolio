<template>
  <div class="descriptions">
    <!-- NDA Icon -->
    <section v-if="projectDetails.orgNDA" class="nda" title="Non Disclosure Agreement Signed">
      <img src="@/assets/logos/NDA.png" />
    </section>

    <h2>{{ projectDetails.title }}</h2>
    <div class="details">
      <section v-if="projectDetails.timeline.academic !== ''">
        <h4>Timeline: </h4>
          <h5 v-if="!projectDetails.timeline.isOnGoing">{{ projectDetails.timeline.academic }}
          </h5> <br v-if="!projectDetails.timeline.isOnGoing" />

          <h5 v-if="projectDetails.timeline.isOnGoing">
            {{ projectDetails.timeline.academic }} (Still On Going)
          </h5> <br  v-if="projectDetails.timeline.isOnGoing"  />
      </section>

      <section>
        <h4>Description: </h4>
          <p v-html="projectDetails.description"></p> <br />
      </section>

      <section>
        <h4>Project Type: </h4>
          <p>{{ projectDetails.projectType }}</p> <br />
      </section>

      <section>
        <h4>Role: </h4>
          <p>{{ projectDetails.role }}</p> <br />
      </section>

      <section v-if="projectDetails.technologies.length > 0">
        <h4>Technologies: </h4>
          <p v-for="(tech, techKey) in projectDetails.technologies"
            :key="techKey">{{ tech.name }}
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
    technologies: {
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
    orgNDA: {
      type: Boolean,
    },
    timeline: {
      type: Object,
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
      technologies: props.technologies,
      projectType: props.projectType,
      role: props.role,
      title: props.title,
      orgName: props.orgName,
      orgExtraLink: props.orgExtraLink,
      orgNDA: props.orgNDA,
      timeline: props.timeline,
    });
    // Mobile Detection
    const isMobile = ref(props.mobile);
    // Truncate String To 50 characters If In Mobile Mode
    if (isMobile.value) projectDetails.description = truncate(projectDetails.description, 50);
    else projectDetails.description = truncate(projectDetails.description, 200);

    watch(() => [
      props.description, props.githubURL, props.liveURL, props.technologies,
      props.projectType, props.role, props.title, props.orgName, props.orgExtraLink, props.orgNDA,
      props.timeline, props.mobile,
    ],
    ([
      desc, githubURL, liveURL, technologies,
      projectType, role, title, orgName, orgExtraLink, orgNDA,
      timeline, mobile,
    ]) => {
      projectDetails.description = desc;
      projectDetails.githubURL = githubURL;
      projectDetails.liveURL = liveURL;
      projectDetails.technologies = technologies;
      projectDetails.projectType = projectType;
      projectDetails.role = role;
      projectDetails.title = title;
      projectDetails.orgName = orgName;
      projectDetails.orgExtraLink = orgExtraLink;
      projectDetails.orgNDA = orgNDA;
      projectDetails.timeline = timeline;

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

  .nda {
    position: absolute;
    top: -5%;
    left: 0;
    width: 4%;
    margin: 5px;

    img {
      width: 100%;
    }
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

    .nda {
      top: -5% !important;
      width: 7% !important;
    }
  }
}
@media screen and (max-width: 600px) {
  .descriptions {
    font-size: 12px;
    overflow-y: auto;
    height: 250px;

    .nda {
      top: -5% !important;
      width: 7% !important;
    }
  }

  /* width of the entire scrollbar */
  ::-webkit-scrollbar,
  ::-webkit-scrollbar { /* Chrome/Opera/Safari */
      width: 4px;
  }
  ::-moz-scrollbar,
  ::-webkit-scrollbar { /* Firefox 19+ */
      width: 4px;
  }

  /* Track */
  ::-webkit-scrollbar-track {
    background-color: rgb(214, 105, 105);
  }

  /* Handle */
  ::-webkit-scrollbar-thumb {
    background-color: rgb(204, 46, 46);
    border-radius: 2px;
  }

  /* Handle on hover */
  ::-webkit-scrollbar-thumb:hover {
    background-color: rgb(204, 46, 46);;
  }
}
@media screen and (max-width: 460px) {
  .descriptions {
    font-size: 9px;
    overflow-y: auto;
    height: 120px;

    .nda {
      top: -5% !important;
      width: 7% !important;
    }
  }

  /* width of the entire scrollbar */
  ::-webkit-scrollbar,
  ::-webkit-scrollbar { /* Chrome/Opera/Safari */
      width: 4px;
  }
  ::-moz-scrollbar,
  ::-webkit-scrollbar { /* Firefox 19+ */
      width: 4px;
  }

  /* Track */
  ::-webkit-scrollbar-track {
    background-color: rgb(214, 105, 105);
  }

  /* Handle */
  ::-webkit-scrollbar-thumb {
    background-color: rgb(204, 46, 46);
    border-radius: 2px;
  }

  /* Handle on hover */
  ::-webkit-scrollbar-thumb:hover {
    background-color: rgb(204, 46, 46);;
  }
}
</style>
